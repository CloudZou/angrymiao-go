package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"angrymiao-go/app/admin/manager/conf"
	"angrymiao-go/app/admin/manager/server/grpc"
	"angrymiao-go/app/admin/manager/server/http"
	"angrymiao-go/app/admin/manager/service"
	"angrymiao-go/punk/log"
	"angrymiao-go/punk/net/trace"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// init log
	log.Init(conf.Conf.Log)
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()
	defer func() {
		log.Close()
		// wait for a while to guarantee that all log messages are written
		time.Sleep(10 * time.Millisecond)
	}()
	// service init
	svc := service.New(conf.Conf)
	grpcSvc := grpc.New(conf.Conf.WardenServer, svc)
	http.Init(conf.Conf, svc)
	log.Info("manager-admin start")
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("manager-admin get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			grpcSvc.Shutdown(context.Background())
			svc.Close()
			log.Info("manager-admin exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
