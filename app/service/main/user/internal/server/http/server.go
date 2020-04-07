package http

import (
	pb "angrymiao-go/app/service/main/user/api"
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/internal/model"
	bm "angrymiao-go/punk/net/http/blademaster"
)

var svc pb.GreeterServer

// New new a bm server.
func New(s pb.GreeterServer) (engine *bm.Engine, err error) {
	svc = s
	engine = bm.DefaultServer(conf.Conf.HTTPServer)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	g := e.Group("/api/v1/user")
	{
		g.GET("/start", howToStart)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Punk{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
