package service

import (
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/punk/log"
	"context"
	"errors"
)


func (s *Service)QQLogin(c context.Context, wechatInfoReq *model.WechatInfoReq) (user *model.User, err error){
	qqTokenResponse, err := s.dao.QQ.GetQQAccessToken(wechatInfoReq.Code, wechatInfoReq.RedirectUrl)
	if err != nil{
		log.Error("s.dao.QQ.GetQQAccessToken(%v,%v) err(%v)", wechatInfoReq.Code, wechatInfoReq.RedirectUrl, err)
		return
	}
	if err = checkErrInfo(qqTokenResponse); err != nil {
		log.Error("checkErrInfo(%v), err(%v)", qqTokenResponse, err)
		return
	}
	qqOpenIdResponse, err := s.dao.QQ.GetQQOpenId(qqTokenResponse.AccessToken)
	if err != nil {
		log.Error("s.dao.QQ.GetQQOpenId(%v) err(%v)", qqTokenResponse.AccessToken, err)
		return
	}
	if err = checkErrInfo(qqTokenResponse); err != nil {
		log.Error("checkErrInfo(%v), err(%v)", qqTokenResponse, err)
		return
	}
	qqUserInfo, err := s.dao.QQ.GetQQUserInfo(qqTokenResponse.AccessToken, qqOpenIdResponse.OpenID)
	if err != nil {
		log.Error("s.dao.QQ.GetQQUserInfo(%v,%v) err(%v)", qqTokenResponse.AccessToken, qqOpenIdResponse.OpenID, err)
		return
	}
	user, err = s.dao.GetUserByQQInfo(qqOpenIdResponse.UnionID)
	if err != nil {
		log.Error("s.dao.GetUserByQQInfo(%v)", qqOpenIdResponse.UnionID, err)
		return
	}
	//user not exist, create
	if user == nil {
		user, err = s.dao.CreateUserWithQQ(qqOpenIdResponse, qqUserInfo)
		if err != nil {
			log.Error("s.dao.CreateUserWithQQ(%v, %v) err(%v)", qqOpenIdResponse, qqUserInfo, err)
			return
		}
	}
	return
}

func checkErrInfo(qqErrorInfo interface{}) (err error){
	switch qqErrorInfo.(type) {
	case model.QQTokenResponse:
		err = errors.New(qqErrorInfo.(model.QQTokenResponse).ErrorDescription)
		return
	case model.QQOpenIDResponse:
		err = errors.New(qqErrorInfo.(model.QQOpenIDResponse).ErrorDescription)
		return
	}
	return
}