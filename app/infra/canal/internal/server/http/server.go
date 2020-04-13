package http

import (
	"angrymiao-go/app/infra/canal/conf"
	"angrymiao-go/app/infra/canal/internal/service"
	bm "angrymiao-go/punk/net/http/blademaster"
)

var (
	cs *service.Canal
)

// New new a bm server.
func New(s *service.Canal) (engine *bm.Engine, err error) {
	cs = s
	engine = bm.DefaultServer(conf.Conf.BM)

	initRouter(engine)
	err = engine.Start()
	return
}

// initRouter init outer router api path.
func initRouter(e *bm.Engine) {
	// init api
	group := e.Group("/x/internal/canal")
	{
		group.GET("/infoc/post", infocPost)
		group.GET("/infoc/current", infocCurrent)
		group.GET("/errors", errors)
		group.POST("/master/check", checkMaster)
		group.POST("/test/sync", syncPos)
	}
}
