package service

import (
	"context"
	"time"

	"angrymiao-go/app/admin/manager/model"
	"angrymiao-go/punk/ecode"
	"angrymiao-go/punk/log"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Auth return user's auth infomation.
func (s *Service) Auth(c context.Context, username string) (res *model.Auth, err error) {
	if err = s.dao.DB().Where("username = ?", username).First(&user).Error; err != nil {
		if err == ecode.NothingFound {
			err = ecode.Int(10001)
			return
		}
		err = errors.Wrapf(err, "s.dao.DB().user.fitst(%s)", username)
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
	res = new(model.Permissions)
	//根据username 招到userid
	res.UID = user.ID
	res.Perms = make([]string, 10)
	res.Admin = true
	return
}


