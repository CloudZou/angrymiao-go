package service

import (
	smsgrpc "angrymiao-go/app/service/main/sms/api"
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/dao"
	"angrymiao-go/punk/log"
)

// Service service.
type Service struct {
	dao       *dao.Dao
	smsgrpc smsgrpc.SmsClient
}

// New new a service and return.
func New(c *conf.Config) (s *Service, cf func(), err error) {
	d, cf, err := dao.New(c)
	if err != nil {
		log.Error("dao.New(%v) err(%v)", c, err)
		return
	}
	smsClient, err := smsgrpc.NewClient(c.SMSClient)
	if err != nil {
		log.Error("smsgrpc.NewClient(%v) err(%v)", c.SMSClient, err)
		return
	}
	s = &Service{
		dao:       d,
		smsgrpc: smsClient,
	}
	cf = s.Close
	return
}
// Close close the resource.
func (s *Service) Close() {
}
