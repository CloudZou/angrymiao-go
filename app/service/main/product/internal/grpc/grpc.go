package grpc

import (
	v1 "angrymiao-go/app/service/main/user/api"
	"context"
	"angrymiao-go/punk/net/rpc/warden"
)

//const
//	//RPC Method
//	//_userInfo              = "RPC.UserInfo"
//)
type Service struct {
	userClient v1.UserInfoClient
}

// New new rpc service.
func NewGrpc(c *warden.ClientConfig) (s *Service) {
	s = &Service{}
	var err error
	//GRPC调用模式，从服务注册中discovery.angrymiao.com找到grpc server 地址
	s.userClient, err = v1.NewClient(nil)
	if err != nil {
		panic(err)
	}
	return
}

func (s *Service) Ping(ctx context.Context) (err error) {
	//_, err = s.demoClient.Ping(ctx, nil)
	return
}

