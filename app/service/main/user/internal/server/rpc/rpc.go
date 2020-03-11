package rpc

import (
	"angrymiao-go/app/service/main/user/api"
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/internal/service"
	"github.com/CloudZou/punk/pkg/log"
	"github.com/CloudZou/punk/pkg/net/rpc"
	"github.com/CloudZou/punk/pkg/net/rpc/context"
)

// RPC server
type RPC struct {
	s *service.Service
}

// New new rpc server.
func New(s *service.Service) (svr *rpc.Server) {
	r := &RPC{s: s}
	svr = rpc.NewServer(conf.Conf.RPCServer)
	if err := svr.Register(r); err != nil {
		panic(err)
	}
	return
}

func (r *RPC) SayHello(c context.Context, arg *string, res *struct{}) (err error) {
	log.Error("Say Hello From RPC Client")
	_, err = r.s.SayHello(c, &api.HelloRequest{Name: *arg})
	return
}
