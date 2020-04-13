package di

import (
	"context"
	"time"

	"angrymiao-go/app/infra/databus/internal/service"

	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
)

//go:generate punk tool wire
type App struct {
	svc *service.Service
	http *bm.Engine
}

func NewApp(svc *service.Service, h *bm.Engine) (app *App, closeFunc func(), err error){
	app = &App{
		svc: svc,
		http: h,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		if err := h.Shutdown(ctx); err != nil {
			log.Error("httpSrv.Shutdown error(%v)", err)
		}
		cancel()
	}
	return
}
