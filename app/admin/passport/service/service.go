package service

import (
	"angrymiao-go/app/admin/passport/conf"
	"angrymiao-go/app/admin/passport/dao"
	"angrymiao-go/punk/log"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

type Service struct {
	dao *dao.Dao
}

// New new a service and return.
func New(c *conf.Config) (s *Service, cf func(), err error) {
	dao, cf, err := dao.NewDao(c)
	if err != nil {
		log.Error("NewDB(%v) err(%v)", c, err)
	}
	s = &Service{
		dao: dao,
	}
	return
}


// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
