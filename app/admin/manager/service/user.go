package service

import (
	"angrymiao-go/app/admin/manager/model"
	"angrymiao-go/punk/ecode"
	"context"
	"github.com/pkg/errors"
)

// Auth return user's auth infomation.
func (s *Service) Auth(c context.Context, username string) (res *model.Auth, err error) {
	user := &model.User{}
	if err = s.dao.DB().Where("username = ?", username).First(&user).Error; err != nil {
		if err == ecode.NothingFound {
			err = ecode.Int(10001)
			return
		}
		err = errors.Wrapf(err, "s.dao.DB().user.first(%s)", username)
		return
	}
	res = &model.Auth{
		UID:      user.ID,
		Username: user.Username,
	}
	res.Perms = make([]string, 10)
	res.Admin = true
	return
}

// Permissions return user's permissions.
func (s *Service) Permissions(c context.Context, username string) (res *model.Permissions, err error) {
	user := &model.User{}
	res = new(model.Permissions)
	//根据username 招到userid
	res.UID = user.ID
	res.Perms = make([]string, 10)
	res.Admin = true
	return
}


