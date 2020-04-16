package dao

import (
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/punk/log"
	"github.com/jinzhu/gorm"
	"github.com/silenceper/wechat/oauth"
)


func (d *Dao)GetUserById(id int64) (user *model.User, err error) {
	err = d.db.Where("id = ? and status = 0", id).Find(user).Error
	if err == gorm.ErrRecordNotFound {
		user = nil
		err = nil
		return
	}
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", id, err)
		return
	}

	return
}


func (d *Dao) GetUserByPhone(phone string) (user *model.User, err error) {
	err = d.db.Where("phone = ? and status = 0", phone).Find(user).Error
	if err == gorm.ErrRecordNotFound {
		user = nil
		err = nil
		return
	}
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", phone, err)
		return
	}

	return
}

func (d *Dao) GetUserByWechatInfo(unionId string) (user *model.User, err error) {
	err = d.db.Where("union_id = ? and auth_type = 'Wechat' and status = 0", unionId).Find(user).Error
	if err == gorm.ErrRecordNotFound {
		user = nil
		err = nil
		return
	}
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", unionId, err)
		return
	}

	return
}

func (d *Dao) GetUserByQQInfo(unionId string) (user *model.User, err error) {
	err = d.db.Where("union_id = ? and auth_type = 'QQ' and status = 0", unionId).Find(user).Error
	if err == gorm.ErrRecordNotFound {
		user = nil
		err = nil
		return
	}
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", unionId, err)
		return
	}

	return
}

func (d *Dao) GetUserByWxOpenId(openId string) (user *model.User, err error) {
	err = d.db.Where("wx_open_id = ? and status = 0", openId).Find(user).Error
	if err == gorm.ErrRecordNotFound {
		user = nil
		err = nil
		return
	}
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", openId, err)
		return
	}
	return
}

func (d *Dao) CreateUserWithPhone(phone string) (user *model.User, err error) {
	user = &model.User{}
	user.Phone = phone
	user.Username = phone
	user.AuthType = model.Mobile
	if d.db.NewRecord(*user) {
		err = d.db.Create(user).Error
		if err != nil {
			log.Error("d.db.Create(%v) err(%v)", *user, err)
			return
		}
	}
	return
}

func (d *Dao)CreateUserWithWechat(oauthUserInfo oauth.UserInfo) (user *model.User, err error) {
	user = &model.User{}
	user.UnionId = oauthUserInfo.Unionid
	user.WxOpenId = oauthUserInfo.OpenID
	user.Username = oauthUserInfo.Nickname
	user.Gender = string(oauthUserInfo.Sex)
	user.HeadPicurl = oauthUserInfo.HeadImgURL
	user.AuthType = model.Wechat
	if d.db.NewRecord(*user) {
		err = d.db.Create(user).Error
		if err != nil {
			log.Error("d.db.Create(%v) err(%v)", *user, err)
			return
		}
	}
	return
}

func (d *Dao)CreateUserWithQQ(qqOpenID model.QQOpenIDResponse, qqUserInfo model.QQUserInfoResponse) (user *model.User, err error) {
	user = &model.User{}
	user.UnionId = qqOpenID.UnionID
	user.WxOpenId = qqOpenID.OpenID
	user.Username = qqUserInfo.Nickname
	user.Gender = qqUserInfo.Gender
	user.HeadPicurl = qqUserInfo.FigureUrl_qq_1
	user.AuthType = model.QQ
	if d.db.NewRecord(*user) {
		err = d.db.Create(user).Error
		if err != nil {
			log.Error("d.db.Create(%v) err(%v)", *user, err)
			return
		}
	}
	return
}

func (d *Dao)UpdateWxOpenIdById(user model.User, wxOpenId string) (err error) {
	err = d.db.Model(&user).Update("wx_open_id", wxOpenId).Error
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", user.ID, err)
		return
	}
	return
}