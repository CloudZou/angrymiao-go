package http

import (
	"net/http"

	"angrymiao-go/app/admin/manager/conf"
	"angrymiao-go/app/admin/manager/service"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/http/blademaster/middleware/permit"
	"angrymiao-go/punk/net/http/blademaster/middleware/verify"
	"angrymiao-go/punk/net/rpc/warden"
)

var (
	mngSvc    *service.Service
	vfySvc    *verify.Verify
	permitSvc *permit.Permit
)

// Init init http sever instance.
func Init(c *conf.Config, s *service.Service) {
	mngSvc = s
	vfySvc = verify.New(nil)
	wardenConf := &warden.ClientConfig{
		NonBlock: true,
	}
	permitSvc = permit.New2(wardenConf)
	// init inner router
	engine := bm.DefaultServer(c.HTTPServer)
	innerRouter(engine)
	// init inner server
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
}

func innerRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/x/admin/manager")
	{
		g.GET("/auth", permitSvc.Verify(), authUser)
		g.GET("/logout", logout)
		g.GET("/permission", vfySvc.Verify, permissions)
	}
}

// ping check server ok.
func ping(c *bm.Context) {
	var err error
	if err = mngSvc.Ping(c); err != nil {
		log.Error("service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
