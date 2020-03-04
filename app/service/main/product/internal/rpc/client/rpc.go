package client

import (
	v1 "angrymiao-go/app/service/main/user/api"
	"context"
	"github.com/CloudZou/punk/pkg/net/rpc"
)

//const
//	//RPC Method
//	//_userInfo              = "RPC.UserInfo"
//)

const (
	_appid = "user.service"
)

var (
	_noArg = &struct{}{}
)

// Service is a question service.
type Service struct {
	client     *rpc.Client2
	demoClient v1.DemoClient
}

// New new rpc service.
func New(c *rpc.ClientConfig) (s *Service) {
	s = &Service{}
	//GRPC调用模式，从服务注册中discovery.angrymiao.com找到grpc server 地址
	//s.demoClient, err = v1.NewClient(nil)
	//if err != nil {
	//	panic(err)
	//}

	s.client = rpc.NewDiscoveryCli(_appid, c)
	return
}

func (s *Service) Ping(ctx context.Context) (err error) {
	err = s.client.Call(ctx, "RPC.Ping", nil, nil)
	//_, err = s.demoClient.Ping(ctx, nil)
	return
}

func (s *Service) SayHello(ctx context.Context, name string) (err error) {
	err = s.client.Call(ctx, "RPC.SayHello", name, nil)
	//_, err = s.demoClient.SayHello(ctx, &v1.HelloReq{Name:name})
	return
}

func (s *Service) SayHelloURL(ctx context.Context, name string) (err error) {
	err = s.client.Call(ctx, "RPC.SayHelloURL", name, nil)
	//_, err = s.demoClient.SayHelloURL(ctx, &v1.HelloReq{Name:name})
	return
}
