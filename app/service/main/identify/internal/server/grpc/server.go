package grpc

import (
	pb "angrymiao-go/app\service\main/identify/api"
	v1 "angrymiao-go/app/service/main/identify/api"
	"angrymiao-go/app/service/main/identify/internal/service"
	"angrymiao-go/punk/ecode"
	"angrymiao-go/punk/log"
	"angrymiao-go/punk/net/metadata"
	"context"
	"fmt"
	"google.golang.org/grpc"

	"angrymiao-go/punk/conf/paladin"
	"angrymiao-go/punk/net/rpc/warden"
)

// New new a grpc server.
func New(cfg *warden.ServerConfig, s *service.Service) *warden.Server {
	w := warden.NewServer(cfg)
	w.Use(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if resp, err = handler(ctx, req); err != nil {
			log.Infov(ctx,
				log.KV("path", info.FullMethod),
				log.KV("caller", metadata.String(ctx, metadata.Caller)),
				log.KV("args", fmt.Sprintf("%v", req)),
				log.KV("args", fmt.Sprintf("%+v", err)))
		}
		return
	})
	v1.RegisterIdentifyServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

var _ v1.IdentifyServer = &server{}
var (
	emptyTokenReply = &v1.GetTokenInfoReply{
		IsLogin: false,
	}
)


func (s *server) GetTokenInfo(ctx context.Context, req *v1.GetTokenInfoReq) (*v1.GetTokenInfoReply, error) {
	token := &v1.GetTokenInfoReq{
		Token: req.Token,
	}
	res, err := s.svr.GetTokenInfo(ctx, token)
	if err != nil {
		if err == ecode.NoLogin {
			return emptyTokenReply, nil
		}
		return nil, err
	}
	return &v1.GetTokenInfoReply{
		IsLogin: true,
		Mid:     res.Mid,
		Expires: res.Expires,
		Csrf:    res.Csrf,
	}, nil
}
