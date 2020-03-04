package grpc

import (
	pb "angrymiao-go/app/service/main/user/api"
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/internal/service"
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/CloudZou/punk/pkg/net/rpc/warden"
)

// New new a grpc rpc.
func New(s *service.Service) (ws *warden.Server, err error) {
	ws = warden.NewServer(conf.Conf.GRPCServer)
	pb.RegisterDemoServer(ws.Server(), &rpc{as: s})
	if ws, err = ws.Start(); err != nil {
		return
	}
	return
}

type rpc struct {
	as *service.Service
}

var _ pb.DemoServer = &rpc{}

func (s *rpc) Ping(ctx context.Context, req *empty.Empty) (reply *empty.Empty, err error) {
	reply, err = s.as.Ping(ctx, req)
	return nil, nil
}

func (s *rpc) SayHello(ctx context.Context, req *pb.HelloReq) (reply *empty.Empty, err error) {
	reply, err = s.as.SayHello(ctx, req)
	return
}

func (s *rpc) SayHelloURL(ctx context.Context, req *pb.HelloReq) (resp *pb.HelloResp, err error) {
	resp, err = s.as.SayHelloURL(ctx, req)
	return
}
