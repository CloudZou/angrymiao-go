package service

import (
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/punk/log"
	"context"
)

func (s *Service)PhoneLogin(c context.Context, req model.PhoneLoginReq) (user *model.User, err error) {
	if err = s.Authenticate(c, req); err != nil {
		log.Error("s.Authenticate(c,%v) err(%v)", req, err)
		return
	}
	user, err = s.dao.GetUserByPhone(req.Phone)
	if err != nil {
		log.Error("s.dao.GetUserByPhone(%v)", req.Phone, err)
		return
	}
	//user not exist, create
	if user == nil {
		user, err = s.dao.CreateUserWithPhone(req.Phone)
		if err != nil {
			log.Error("s.dao.CreateUserWithPhone(%v) err(%v)", req.Phone, err)
			return
		}
	}
	return
}
