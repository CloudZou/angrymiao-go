package http

import (
	"angrymiao-go/app/service/main/product/conf"
	"angrymiao-go/app/service/main/product/internal/grpc"
	"angrymiao-go/app/service/main/product/internal/rpc"
	"angrymiao-go/app/service/main/product/internal/service"
	"net/http"

	"github.com/CloudZou/punk/pkg/log"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
)

var svc *service.Service
var rpcSvc *rpc.Service
var grpcSvc *grpc.Service

func New(s *service.Service) (engine *bm.Engine, err error) {
	svc = s
	rpcSvc = rpc.New(conf.Conf.RPCClient)
	grpcSvc = grpc.NewGrpc(conf.Conf.UserGrpcClient)
	engine = bm.DefaultServer(conf.Conf.BM)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	g := e.Group("/api/v1")
	{
		g.GET("/sayHello", sayHello)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func sayHello(c *bm.Context) {
	err := grpcSvc.SayHello(c, c.Params.ByName("name"))
	if err != nil {
		log.Error("say hello error(%v)", err)
	}
	c.JSON(nil, nil)
	return
}

func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}
