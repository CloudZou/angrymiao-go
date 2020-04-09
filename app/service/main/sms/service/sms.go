package service

import (
	pb "angrymiao-go/app/service/main/sms/api"
	"angrymiao-go/punk/ecode"
	"context"
	"regexp"
)
const (
	_regionId = "cn-hangzhou"
	_signName = "怒喵科技"
	mobilePattern = "^((13[0-9])|(14[1,4,5,6,7,8])|(15[^4])|(16[5-7])|(17[0-8])|(18[0-9])|(19[1,8,9]))\\d{8}$"
)

var mobileReg, _ = regexp.Compile(mobilePattern)

func (s *Service) Send(c context.Context, req *pb.SendReq)  (res *pb.SendReply, err error) {
	if req.Mobile == "" {
		err = ecode.RequestErr
		return
	}

	if match := mobileReg.MatchString(req.Mobile); !match {
		err = ecode.SmsMobilePatternErr
		return
	}

	req.SignName = _signName
	req.RegionId = _regionId
	err = s.dao.PubSingle(c, req)
	return
}