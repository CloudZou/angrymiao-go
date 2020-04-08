package di

import (
	"angrymiao-go/app/service/main/identify/internal/service"
	"context"
	"time"

	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/rpc/warden"
)

//go:generate punk tool wire
type App struct {
	svc *service.Service
	http *bm.Engine
	grpc *warden.Server
}

func NewApp(svc *service.Service, h *bm.Engine, g *warden.Server) (app *App, closeFunc func(), err error){
	app = &App{
		svc: svc,
		http: h,
		grpc: g,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		if err := g.Shutdown(ctx); err != nil {
			log.Error("grpcSrv.Shutdown error(%v)", err)
		}
		if err := h.Shutdown(ctx); err != nil {
			log.Error("httpSrv.Shutdown error(%v)", err)
		}
		cancel()
	}
	return
}
