package main

import (
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/server/http"
	"angrymiao-go/app/service/main/user/service"
	"context"
	"flag"
	"angrymiao-go/punk/conf/env"
	"angrymiao-go/punk/naming"
	"angrymiao-go/punk/naming/discovery"
	"os"
	"os/signal"
	"syscall"
	"time"

	"angrymiao-go/punk/log"
	xip "angrymiao-go/punk/net/ip"
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

	var cancel context.CancelFunc
	if env.IP == "" {
		ip := xip.InternalIP()
		hn, _ := os.Hostname()
		env.Region = "region01"
		env.Zone = "zone01"
		dis := discovery.New(nil)
		ins := &naming.Instance{
			AppID:    "user.service",
			Hostname: hn,
			Addrs: []string{
				"http://" + ip + ":" + env.HTTPPort,
				"gorpc://" + ip + ":" + env.GORPCPort,
				"grpc://" + ip + ":" + env.GRPCPort,
				"http://test.angrymiao.com",
			},
		}
		if cancel, err = dis.Register(context.Background(), ins); err != nil {
			panic(err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if cancel != nil {
				cancel()
			}
			closeFunc()
			log.Info("user exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
