package http

import (
	"angrymiao-go/app/service/main/sms/conf"
	"angrymiao-go/app/service/main/sms/service"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/http/blademaster/middleware/verify"
)


var (
	smsSvc *service.Service
	idfSvc *verify.Verify
)

// Init init http sever instance.
func Init(c *conf.Config, s *service.Service) {
	idfSvc = verify.New(c.Verify)
	smsSvc = s
	engine := bm.DefaultServer(c.HTTPServer)
	route(engine)
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
}

func route(e *bm.Engine) {
	e.Register(register)
	g := e.Group("/x/internal/sms", bm.CORS([]string{"*"}))
	{
		g.POST("/send", idfSvc.Verify, send)
	}
}


func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}
