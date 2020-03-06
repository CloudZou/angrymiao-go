package rpc

import (
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
	client *rpc.Client2
}

// New new rpc service.
func New(c *rpc.ClientConfig) (s *Service) {
	s = &Service{}
	s.client = rpc.NewDiscoveryCli(_appid, c)
	return
}

func (s *Service) Ping(ctx context.Context) (err error) {
	err = s.client.Call(ctx, "RPC.Ping", nil, nil)
	return
}

func (s *Service) SayHello(ctx context.Context, name string) (err error) {
	err = s.client.Call(ctx, "RPC.SayHello", name, nil)
	return
}

func (s *Service) SayHelloURL(ctx context.Context, name string) (err error) {
	err = s.client.Call(ctx, "RPC.SayHelloURL", name, nil)
	return
}
