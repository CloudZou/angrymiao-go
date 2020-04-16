package grpc

import (
	"angrymiao-go/app/admin/manager/service"
	"angrymiao-go/punk/net/rpc/warden"
)

// server .
type server struct {
	srv *service.Service
}

// New return warden server.
func New(cfg *warden.ServerConfig, srv *service.Service) *warden.Server {
	w := warden.NewServer(cfg)
	//pb.RegisterPermitServer(w.Server(), &server{srv: srv})
	var err error
	if w, err = w.Start(); err != nil {
		panic(err)
	}
	return w
}