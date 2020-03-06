package main

import (
	"angrymiao-go/app/service/main/product/conf"
	"flag"
	"github.com/CloudZou/punk/pkg/net/trace"
	"os"
	"os/signal"
	"syscall"
	"time"

	"angrymiao-go/app/service/main/product/internal/di"
	"github.com/CloudZou/punk/pkg/log"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("product start")
	trace.Init(nil)
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("product exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
