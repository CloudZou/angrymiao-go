package service

import (
	"angrymiao-go/app/job/sms/dao"
	smsgrpc "angrymiao-go/app/service/main/sms/api"
	smsmdl "angrymiao-go/app/service/main/sms/model"
	"angrymiao-go/punk/conf/env"
	"angrymiao-go/punk/log"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"regexp"
	"time"
)

var _contentRe = regexp.MustCompile(`\d`)

func (s *Service) subproc() {
	defer s.waiter.Done()
	for {
		item, ok := <-s.databus.Messages()
		if !ok {
			close(s.sms)
			log.Info("databus: sms-job subproc consumer exit!")
			return
		}
		s.smsCount++
		msg := new(smsgrpc.SendReq)
		if err := json.Unmarshal(item.Value, &msg); err != nil {
			log.Error("json.Unmarshal (%v) error(%v)", string(item.Value), err)
			continue
		}
		log.Info("subproc topic(%s) key(%s) partition(%v) offset(%d) message(%s)", item.Topic, item.Key, item.Partition, item.Offset, item.Value)
		// 黑名单，用于压测
		if _, ok := s.blacklist[msg.Country+msg.Mobile]; ok {
			log.Info("country(%s) mobile(%s) in blacklist", msg.Country, msg.Mobile)
			item.Commit()
			continue
		}
		s.sms <- msg
		item.Commit()
	}
}

func (s *Service) smsproc() {
	defer s.waiter.Done()
	var (
		msgid string
	)
	for {
		m, ok := <-s.sms
		if !ok {
			log.Info("smsproc exit!")
			return
		}
		if m.Mobile == "" {
			log.Error("invalid country or mobile, info(%+v)", m)
			continue
		}
		smsRequest := s.newSmsRequest(m)

		content := _contentRe.ReplaceAllString(string(smsRequest.Content), "*")
		l := &smsmdl.ModelUserActionLog{Mobile: m.Mobile, Content: content, Type: smsmdl.TypeSms, Action: smsmdl.UserActionTypeSend}

		s.smsp.Lock()
		s.smsp.Ring = s.smsp.Next()
		s.smsp.Unlock()
		response, err := s.smsClient.SendSms(smsRequest)
		msgid = response.BizId
		l.MsgID = msgid
		if err == nil {
			l.Status = smsmdl.UserActionSendSuccessStatus
			l.Desc = smsmdl.UserActionSendSuccessDesc
			dao.PromInfo(fmt.Sprintf("service:success"))
			log.Info("send sms(%v) success", m)
		} else {
			l.Status = smsmdl.UserActionSendFailedStatus
			l.Desc = smsmdl.UserActionSendFailedDesc
			dao.PromError("service:sms")
			log.Error("send sms(%v) error(%v)", m, err)
			s.cache.Do(context.Background(), func(ctx context.Context) {
				s.dao.SendWechat(fmt.Sprintf("sms-job send msg(%d) error(%v)", m.Mobile, err))
			})

		}
		l.MsgID = msgid
		l.Ts = time.Now().Unix()
		s.sendUserActionLog(l)
	}
}

func (s *Service) newSmsRequest(sendReq *smsgrpc.SendReq) (request *dysmsapi.SendSmsRequest){
	request = dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = sendReq.Mobile
	request.SignName = sendReq.SignName
	request.TemplateCode = sendReq.Tcode
	request.TemplateParam = sendReq.Tparam
	return
}

func (s *Service) monitorproc() {
	if env.DeployEnv != env.DeployEnvProd {
		return
	}
	var smsCount int64
	for {
		time.Sleep(time.Duration(s.c.Sms.MonitorProcDuration))
		if s.smsCount-smsCount == 0 {
			msg := fmt.Sprintf("sms-job sms did not consume within %s seconds", time.Duration(s.c.Sms.MonitorProcDuration).String())
			s.dao.SendWechat(msg)
			log.Warn(msg)
		}
		smsCount = s.smsCount
	}
}
