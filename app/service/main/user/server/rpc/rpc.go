package rpc

import (
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/service"
	"angrymiao-go/punk/net/rpc"
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

