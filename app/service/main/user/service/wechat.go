package service

import (
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/punk/log"
	"context"
)

func (s *Service)WechatLogin(c context.Context, wechatInfoReq *model.WechatInfoReq) (user *model.User, err error){
	oauthUserInfo, err := s.dao.OAuth2GetUserInfo(wechatInfoReq.Code)
	if err != nil {
		log.Error("s.dao.OAuth2GetUserInfo(%v) err(%v)", wechatInfoReq.Code, err)
		return
	}
	user, err = s.dao.GetUserByWechatInfo(oauthUserInfo.Unionid)
	if err != nil {
		log.Error("s.dao.GetUserByWechatInfo(%v)", oauthUserInfo.Unionid, err)
		return
	}
	//user not exist, create
	if user == nil {
		user, err = s.dao.CreateUserWithWechat(oauthUserInfo)
		if err != nil {
			log.Error("s.dao.CreateUserWithWechat(%v) err(%v)", oauthUserInfo, err)
			return
		}
	}
	return
}