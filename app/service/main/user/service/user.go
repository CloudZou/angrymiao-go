package service

import (
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/punk/ecode"
	bm "angrymiao-go/punk/net/http/blademaster"
)

func (s *Service)GetCurrentUser(c *bm.Context)(user *model.User, err error) {
	userId, ok := c.Get("mid")
	if !ok {
		err = ecode.NoLogin
		return
	}
	id, ok := userId.(int64)
	if !ok {
		err = ecode.NoLogin
		return
	}
	user, err = s.dao.GetUserById(id)
	if err != nil || user == nil {
		err = ecode.NoLogin
		return
	}
	return
}

func (s *Service)UpdateWxOpenIdById(user model.User, wxOpenId string) (err error) {
	err = s.dao.UpdateWxOpenIdById(user, wxOpenId)
	return
}


func (s *Service) GetUserById(id int64) (user *model.User, err error) {
	user, err = s.dao.GetUserById(id)
	return
}