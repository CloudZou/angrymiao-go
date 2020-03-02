package di

import (
	"context"
	"time"

	"angrymiao-go/app/infra/canal/internal/service"

	"github.com/CloudZou/punk/pkg/log"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
)

//go:generate punk tool wire
type App struct {
	svc  *service.Canal
	http *bm.Engine
}

func NewApp(svc *service.Canal, h *bm.Engine) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:  svc,
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
