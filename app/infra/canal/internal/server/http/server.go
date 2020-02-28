package http

import (
	"angrymiao-go/app/infra/canal/internal/service"
	"net/http"
	"github.com/CloudZou/punk/pkg/conf/paladin"
	"github.com/CloudZou/punk/pkg/log"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
)

var (
	svc *service.Service
	cs *service.Canal
)
// New new a bm server.
func New(s *service.Service, c *service.Canal) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	cs = c
	engine = bm.DefaultServer(&cfg)

	initRouter(engine)
	err = engine.Start()
	return
}


// initRouter init outer router api path.
func initRouter(e *bm.Engine) {
	// init api
	e.Ping(ping)
	group := e.Group("/x/internal/canal")
	{
		group.GET("/infoc/post", infocPost)
		group.GET("/infoc/current", infocCurrent)
		group.GET("/errors", errors)
		group.POST("/master/check", checkMaster)
		group.POST("/test/sync", syncPos)
	}
}


func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
