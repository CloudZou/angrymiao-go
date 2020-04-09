package service

import (
	"angrymiao-go/app/service/main/sms/conf"
	"angrymiao-go/app/service/main/sms/dao"
)

// Service biz service def.
type Service struct {
	c   *conf.Config
	dao *dao.Dao

	userIds       map[string]int64 // users' ids
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return s
}


// Wait wait all closed.
func (s *Service) Wait() {
}

// Close close all dao.
func (s *Service) Close() {
	s.dao.Close()
}
