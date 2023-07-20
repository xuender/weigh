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
}

// NewService creates a new instance of Service.
func NewService(cfg *pb.Config) *Service {
	eng := gin.Default()
	service := &Service{
		eng:    eng,
		client: resty.New(),
		cfg:    cfg,
	}
	service.pool = pools.New(int(cfg.PoolSize), service.Execute)

	eng.GET("/", service.ping)
	eng.POST("/proxy", service.proxy)

	return service
}

func NewHandler(service *Service) http.Handler {
	return service.Handler()
}

func (p *Service) proxy(ctx *gin.Context) {
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
	reqmsg.Response = responses
	ctx.ProtoBuf(http.StatusOK, reqmsg)
}

func (p *Service) Execute(req *pb.Request, num int) *pb.Response {
	request := p.client.R()
	request.Body = req.Body

	for _, item := range req.Head {
		request.Header.Add(item.Name, item.Value)
	}

	res, err := request.Execute(req.Method, req.URL)
	if err != nil {
		logs.E.Println(num, err)

		return &pb.Response{
			Error:      err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
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
