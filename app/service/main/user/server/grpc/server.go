package grpc

import (
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/service"
	"angrymiao-go/punk/net/rpc/warden"
	pb "angrymiao-go/app/service/main/user/api"
	"angrymiao-go/punk/util"
	"context"
)

// New new a grpc rpc.
func New(s *service.Service) (ws *warden.Server, err error) {
	ws = warden.NewServer(conf.Conf.GRPCServer)
	pb.RegisterUserInfoServer(ws.Server(), &server{as: s})
	if ws, err = ws.Start(); err != nil {
		return
	}
	return
}

type server struct {
	as *service.Service
}

func (s *server)GetUserInfoByID(c context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	user, err := s.as.GetUserById(req.UserId)
	if err != nil {
		return nil, err
	}
	var userInfoResp pb.GetUserInfoResp
	util.StructCopy(&userInfoResp, user)
	return &userInfoResp, nil
}

