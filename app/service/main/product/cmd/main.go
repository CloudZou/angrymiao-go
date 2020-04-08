package main

import (
	"angrymiao-go/app/service/main/product/conf"
	"flag"
	"angrymiao-go/punk/conf/env"
	"angrymiao-go/punk/net/trace"
	"os"
	"os/signal"
	"syscall"
	"time"

	"angrymiao-go/app/service/main/product/internal/di"
	"angrymiao-go/punk/log"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("product start")
	env.AppID = "product.service"
	trace.Init(conf.Conf.UdpTraceConfig)
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
