package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-resty/resty/v2"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/pools"
	"github.com/xuender/weigh/pb"
)

// Service is proxy service.
type Service struct {
	eng    *gin.Engine
	client *resty.Client
	pool   *pools.Pool[*pb.Request, *pb.Response]
	cfg    *pb.Config
	limits *Limits
}

// NewService creates a new instance of Service.
func NewService(cfg *pb.Config, limits *Limits) *Service {
	eng := gin.Default()
	service := &Service{
		eng:    eng,
		client: resty.New(),
		cfg:    cfg,
		limits: limits,
	}
	service.pool = pools.New(int(cfg.PoolSize), service.Execute)

	eng.GET("/", service.ping)
	eng.POST("/proxy", service.proxy)
	eng.POST("/api", service.api)

	return service
}

func NewHandler(service *Service) http.Handler {
	return service.Handler()
}

func (p *Service) proxy(ctx *gin.Context) {
	p.run(ctx, true)
}

func (p *Service) api(ctx *gin.Context) {
	p.run(ctx, false)
}

func (p *Service) run(ctx *gin.Context, old bool) {
	// start := time.Now()
	msg := new(pb.Msg)

	if err := ctx.ShouldBindBodyWith(msg, binding.ProtoBuf); err != nil {
		logs.E.Println(err)
		ctx.String(http.StatusBadRequest, err.Error())

		return
	}

	async, serial := p.cfg.Group(msg.Request)
	responses := p.pool.Post(async)
	responses = append(responses, lo.Map(serial, p.Execute)...)
	// endLog(start, msg, responses)

	reqmsg := new(pb.Msg)

	if old {
		for index, res := range responses {
			res.Compatible(msg.Request[index].Id)
		}
	}

	reqmsg.Response = responses
	ctx.ProtoBuf(http.StatusOK, reqmsg)
}

func (p *Service) Execute(req *pb.Request, num int) *pb.Response {
	if err := p.limits.Check(req.URL); err != nil {
		logs.E.Println(num, err)

		return pb.NewErr(err)
	}

	request := p.client.R()
	request.Body = req.Body

	for _, item := range req.Head {
		request.Header.Add(item.Name, item.Value)
	}

	res, err := request.Execute(req.Method, req.URL)
	if err != nil {
		logs.E.Println(num, err)

		return pb.NewErr(err)
	}

	resHeader := res.Header()
	header, index := make([]*pb.Head, len(resHeader)), 0

	for k, v := range resHeader {
		header[index] = &pb.Head{Name: k, Value: v[0]}
		index++
	}

	body := res.Body()
	code := int32(res.StatusCode())

	if code != http.StatusOK {
		logs.D.Println(num, string(body))
	}

	return &pb.Response{
		StatusCode:    code,
		Status:        res.Status(),
		Body:          body,
		Header:        header,
		ContentLength: int64(len(body)),
	}
}

func (p *Service) ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PONG")
}

func (p *Service) Handler() http.Handler {
	return p.eng
}
