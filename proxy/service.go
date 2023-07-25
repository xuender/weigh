package proxy

import (
	"crypto/tls"
	"net/http"
	"time"

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
	eng     *gin.Engine
	client  *resty.Client
	clients map[string]*resty.Client
	pool    *pools.Pool[*pb.Request, *pb.Response]
	cfg     *pb.Config
	limits  *Limits
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
	transport := &http.Transport{
		// nolint: gosec
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        int(cfg.MaxIdleConns),
		MaxIdleConnsPerHost: int(cfg.MaxIdleConnsPerHost),
	}

	service.pool = pools.New(int(cfg.PoolSize), service.Execute)
	service.client.SetTimeout(time.Duration(cfg.TimeoutSecond) * time.Second)
	service.client.SetTransport(transport)

	if len(cfg.Timeout) > 0 {
		service.clients = make(map[string]*resty.Client, len(cfg.Timeout))

		for key, second := range cfg.Timeout {
			client := resty.New()

			client.SetTransport(transport)
			client.SetCloseConnection(true)
			client.SetTimeout(time.Duration(second) * time.Second)
			service.clients[key] = client
		}
	}

	if len(cfg.QPS) > 0 {
		for key, qps := range cfg.QPS {
			service.limits.Add(key, int(qps))
		}
	}

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
	start := time.Now()
	msg := &pb.Msg{}

	if err := ctx.ShouldBindBodyWith(msg, binding.ProtoBuf); err != nil {
		logs.E.Println(err)
		ctx.String(http.StatusBadRequest, err.Error())

		return
	}

	// logs.I.Println(string(lo.Must1(sonic.Marshal(msg))))

	async, serial := p.cfg.Group(msg.Request)
	responses := p.pool.Post(async)
	responses = append(responses, lo.Map(serial, p.Execute)...)

	if dur := time.Since(start); dur > time.Second*10 {
		logs.I.Printf("LONG_TIME: %v size: %d url: %s\n", dur, len(msg.Request), msg.Request[0].Uri)
	}

	if old {
		for index, res := range responses {
			res.Compatible(msg.Request[index].Id)
		}
	}

	ctx.ProtoBuf(http.StatusOK, &pb.Msg{Response: responses})
}

func (p *Service) getClient(url string) *resty.Client {
	if len(p.clients) == 0 {
		return p.client
	}

	for key, client := range p.clients {
		if Has(url, key) {
			return client
		}
	}

	return p.client
}

func (p *Service) Execute(pbreq *pb.Request, num int) *pb.Response {
	if err := p.limits.Check(pbreq.URL); err != nil {
		logs.E.Println(num, err)

		return pb.NewErr(err)
	}

	req := p.getClient(pbreq.URL).R()
	req.Body = pbreq.Body

	for _, item := range pbreq.Head {
		req.Header.Add(item.Name, item.Value)
	}

	res, err := req.Execute(pbreq.Method, pbreq.URL)
	if err != nil {
		logs.E.Printf("%d read ERR: %s [%d] %s %s\n", num, req.Method, res.StatusCode(), req.URL, err.Error())

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
