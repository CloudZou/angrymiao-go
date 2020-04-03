package http

import (
	"angrymiao-go/app/admin/passport/conf"
	"angrymiao-go/app/admin/passport/service"
	"net/http"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(c *conf.Config, s *service.Service) (engine *bm.Engine, err error) {
	engine = bm.DefaultServer(c.Server)
	svc = s
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/x/admin/passport")
	{
		g.POST("/signIn", signIn)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func signIn(c *bm.Context) {
	//signIn,生成cookie返回给前端，并写入session到redis
	c.JSON("", nil)
}