package http

import (
	"angrymiao-go/app/service/main/user/conf"
	"net/http"

	pb "angrymiao-go/app/service/main/user/api"
	"angrymiao-go/app/service/main/user/internal/model"
	"github.com/CloudZou/punk/pkg/log"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
)

var svc pb.DemoServer

// New new a bm server.
func New(s pb.DemoServer) (engine *bm.Engine, err error) {
	svc = s
	engine = bm.DefaultServer(conf.Conf.HTTPServer)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/user")
	{
		g.GET("/start", howToStart)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Punk{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
