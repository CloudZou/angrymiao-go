package grpc

import (
	pb "angrymiao-go/app/service/main/product/api"
	"angrymiao-go/app/service/main/product/conf"

	"github.com/CloudZou/punk/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc pb.DemoServer) (ws *warden.Server, err error) {
	var cfg *warden.ServerConfig = conf.Conf.RPCServer
	ws = warden.NewServer(cfg)
	pb.RegisterDemoServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
