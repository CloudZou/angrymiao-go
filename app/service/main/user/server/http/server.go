package http

import (
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/service"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
)

var srv service.Service

// New new a bm server.
// New new a bm server.
func New(c *conf.Config, s *service.Service) (engine *bm.Engine, err error) {
	srv = s
	engine = bm.DefaultServer(c.BM)
	initRouter(engine)
	err = engine.Start()
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	g := e.Group("/identify")
	{

	}
}
