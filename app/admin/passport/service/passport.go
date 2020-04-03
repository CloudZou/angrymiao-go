package service

import "angrymiao-go/app/admin/passport/model"

func (s *Service)SignIn(signInForm *model.SignInForm) (ret bool, err error) {
	ret, err = s.dao.Auth(signInForm)
	return
}