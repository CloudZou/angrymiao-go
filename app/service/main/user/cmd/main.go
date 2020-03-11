package main

import (
	"angrymiao-go/app/service/main/user/conf"
	"context"
	"flag"
	"github.com/CloudZou/punk/pkg/conf/env"
	"github.com/CloudZou/punk/pkg/naming"
	"github.com/CloudZou/punk/pkg/naming/discovery"
	"github.com/CloudZou/punk/pkg/net/trace"
	"os"
	"os/signal"
	"syscall"
	"time"

	"angrymiao-go/app/service/main/user/internal/di"
	"github.com/CloudZou/punk/pkg/log"
	xip "github.com/CloudZou/punk/pkg/net/ip"
)

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("user start")
	conf.Init()
	env.AppID = "user.service"
	trace.Init(conf.Conf.UdpTraceConfig)
	_, closeFunc, err := di.InitApp()
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
