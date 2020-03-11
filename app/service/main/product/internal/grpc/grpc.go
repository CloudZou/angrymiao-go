package grpc

import (
	v1 "angrymiao-go/app/service/main/user/api"
	"context"
	"github.com/CloudZou/punk/pkg/net/rpc/warden"
)

//const
//	//RPC Method
//	//_userInfo              = "RPC.UserInfo"
//)
type Service struct {
	demoClient v1.GreeterClient
}

// New new rpc service.
func NewGrpc(c *warden.ClientConfig) (s *Service) {
	s = &Service{}
	var err error
	//GRPC调用模式，从服务注册中discovery.angrymiao.com找到grpc server 地址
	s.demoClient, err = v1.NewClient(nil)
	if err != nil {
		panic(err)
	}
	return
}

func (s *Service) Ping(ctx context.Context) (err error) {
	//_, err = s.demoClient.Ping(ctx, nil)
	return
}

func (s *Service) SayHello(ctx context.Context, name string) (err error) {
	_, err = s.demoClient.SayHello(ctx, &v1.HelloRequest{Name: "test"})
	return
}
