package grpc

import (
	pb "angrymiao-go/app/service/main/sms/api"
	"angrymiao-go/app/service/main/sms/service"
	"angrymiao-go/punk/net/rpc/warden"
)

// New return warden server.
func New(cfg *warden.ServerConfig, srv *service.Service) *warden.Server {
	w := warden.NewServer(cfg)
	pb.RegisterSmsServer(w.Server(), srv)
	var err error
	if w, err = w.Start(); err != nil {
		panic(err)
	}
	return w
}