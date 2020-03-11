package service

import (
	"context"
	"fmt"

	pb "angrymiao-go/app/service/main/user/api"
	"angrymiao-go/app/service/main/user/internal/dao"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.GreeterServer), new(*Service)))

// Service service.
type Service struct {
	dao *dao.Dao
}

// New new a service and return.
func New(d *dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		dao: d,
	}
	cf = s.Close
	return
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (reply *pb.HelloReply, err error) {
	reply = &pb.HelloReply{
		Message: "com from server",
		Success: true,
	}
	fmt.Printf("hello %s", req.Name)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
