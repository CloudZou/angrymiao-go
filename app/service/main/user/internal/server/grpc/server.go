package grpc

import (
	"angrymiao-go/app/service/main/user/api"
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/internal/service"
	"context"

	"github.com/CloudZou/punk/pkg/net/rpc/warden"
)

// New new a grpc rpc.
func New(s *service.Service) (ws *warden.Server, err error) {
	ws = warden.NewServer(conf.Conf.GRPCServer)
	api.RegisterGreeterServer(ws.Server(), &rpc{as: s})
	if ws, err = ws.Start(); err != nil {
		return
	}
	return
}

type rpc struct {
	as *service.Service
}

var _ api.GreeterServer = &rpc{}

func (s *rpc) SayHello(ctx context.Context, req *api.HelloRequest) (reply *api.HelloReply, err error) {
	reply, err = s.as.SayHello(ctx, req)
	return
}
