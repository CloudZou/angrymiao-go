package service

import (
	"angrymiao-go/app/job/sms/conf"
	"angrymiao-go/app/job/sms/dao"
	"angrymiao-go/app/job/sms/model"
	smsgrpc "angrymiao-go/app/service/main/sms/api"
	"angrymiao-go/punk/log"
	"angrymiao-go/punk/queue/databus"
	"angrymiao-go/punk/sync/pipeline/fanout"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"sync"
)
const (
	_proLen = 1
)

// Service biz service def.
type Service struct {
	c   *conf.Config
	dao *dao.Dao
	databus *databus.Databus
	smsgrpc smsgrpc.SmsClient
	waiter *sync.WaitGroup
	smsClient *dysmsapi.Client
	sms   chan *smsgrpc.SendReq
	smsp *model.ConcurrentRing
	cache     *fanout.Fanout
	smsCount  int64
	blacklist map[string]struct{}
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	var err error
	client, err := newSmsClient(c.AliyunSmsConfig)
	if err != nil {
		log.Error("newSmsClient(%v) err(%v)", c.AliyunSmsConfig, err)
		return
	}
	s = &Service{
		c:   c,
		dao: dao.New(c),
		databus: databus.New(c.Databus),
		waiter: new(sync.WaitGroup),
		smsClient: client,
		sms: make(chan *smsgrpc.SendReq, 10240),
		smsp: model.NewConcurrentRing(_proLen),
		cache:     fanout.New("async-task", fanout.Worker(1), fanout.Buffer(10240)),
	}
	s.initBlacklist()
	if s.smsgrpc, err = smsgrpc.NewClient(c.SmsGRPC); err != nil {
		panic(err)
	}
	s.waiter.Add(1)
	go s.subproc()
	go s.monitorproc()
	for i := 0; i < s.c.Sms.SingleSendProc; i++ {
		s.waiter.Add(1)
		go s.smsproc()
	}
	return
}

func newSmsClient(config *conf.AliyunSmsConfig) (client *dysmsapi.Client, err error) {
	client, err = dysmsapi.NewClientWithAccessKey(config.RegionID, config.AccessKeyID, config.AccessKeySecret)
	return
}


func (s *Service) initBlacklist() {
	s.blacklist = make(map[string]struct{})
	for _, v := range s.c.Sms.Blacklist {
		s.blacklist[v] = struct{}{}
	}
}

// Close close all dao.
func (s *Service) Close() {
	s.dao.Close()
}
