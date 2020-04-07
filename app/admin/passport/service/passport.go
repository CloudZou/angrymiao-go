package service

import (
	"angrymiao-go/app/admin/passport/model"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
)

func (s *Service)SignIn(ctx *bm.Context, key string, signInForm *model.SignInForm) (err error) {
	err = s.dao.Auth(ctx, key, signInForm)
	if err != nil {
		log.Error("s.dao.Auth(ctx, %v, %v)", key, signInForm)
		return
	}
	return
}