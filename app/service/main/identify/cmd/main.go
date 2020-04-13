package main

import (
	"angrymiao-go/app/service/main/identify/conf"
	"angrymiao-go/app/service/main/identify/internal/server/http"
	"angrymiao-go/app/service/main/identify/internal/service"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"
	"angrymiao-go/punk/log"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("passport start")
	s, closeFunc, err := service.New(conf.Conf)
	http.New(conf.Conf, s)
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
			log.Info("identify exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
