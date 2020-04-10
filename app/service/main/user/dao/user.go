package dao

import (
	"angrymiao-go/app/service/main/user/internal/model"
	"angrymiao-go/punk/log"
)

func (d *Dao) getUserByPhone(phone string) (user model.User, err error) {
	err = d.db.Where("phone = ? and status = 0", phone).Find(&user).Error
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", phone, err)
		return
	}
	return
}

func (d *Dao) getUserByWxOpenId(openId string) (user model.User, err error) {
	err = d.db.Where("wx_open_id = ? and status = 0", openId).Find(&user).Error
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", openId, err)
		return
	}
	return
}
