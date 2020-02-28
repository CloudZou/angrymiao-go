package http

import (
	"angrymiao-go/app/infra/databus/internal/service"
	"angrymiao-go/app/infra/databus/tcp"
	"github.com/CloudZou/punk/pkg/conf/paladin"

	"github.com/CloudZou/punk/pkg/log"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
)


var (
	svc *service.Service
)

// Init init https
func New(s *service.Service) (engine *bm.Engine, err error) {
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
	// router
	engine = bm.DefaultServer(&cfg)
	initRouter(engine)
	// init internal server
	if err := engine.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
	return
}

// initRouter init local router api path.
func initRouter(e *bm.Engine) {
	e.Ping(ping)
	e.Register(register)
	e.GET("/databus/consumer/addrs", consumerAddrs)
	e.POST("/databus/pub", pub)
}

// ping check server ok
func ping(c *bm.Context) {
}

// register provid for discovery.
func register(c *bm.Context) {
	c.JSON(map[string]struct{}{
		"data": struct{}{},
	}, nil)
}

// consumerAddrs get consumer addrs.
func consumerAddrs(c *bm.Context) {
	group := c.Request.Form.Get("group")
	c.JSON(tcp.ConsumerAddrs(group))
}
