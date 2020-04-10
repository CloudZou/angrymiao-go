package service

import (
	smsgrpc "angrymiao-go/app/service/main/sms/api"
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/punk/log"
	"context"
	"fmt"
)

func (s *Service)ValidateAndSend(c context.Context, mobile string, smsType model.SmsType) (err error) {
	frequencyKey := model.KeyOfMobileGetCaptchaFrequency(mobile, smsType)
	if err := s.dao.ValidateFrequency(c, frequencyKey);err != nil {
		log.Error("s.dao.GetFrequency(c,%v) err(%v)", frequencyKey, err)
		return
	}
	captcha := s.dao.GenerateRandomNumericString()
	captchaKey := model.KeyOfMobileCaptcha(mobile, smsType)
	if err := s.dao.AddCaptchaCache(c, captchaKey, captcha); err != nil {
		log.Error("s.dao.AddCaptchaCache(c,%v,%v) err(%v)", captchaKey, captcha, err)
		return
	}
	//动态验证码
	sendReq := &smsgrpc.SendReq{
		Mobile:               mobile,
		SignName:             "怒喵科技",
		Tcode:                "SMS_180950718",
		Tparam:               fmt.Sprintf("{\"code\":%s}", captcha),
	}
	if _, err := s.smsgrpc.Send(c, sendReq); err != nil {
		log.Error("s.smsgrpc.Send(c,%v) err(%v)", sendReq, err)
		return
	}
	return
}