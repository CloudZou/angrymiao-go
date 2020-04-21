package dao

import (
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/punk/log"
)

func (d *Dao)CreateUserAddress(address *model.UserAddress) (err error) {
	if d.db.NewRecord(*address) {
		err = d.db.Create(address).Error
		if err != nil {
			log.Error("d.db.Create(%v) err(%v)", address, err)
			return
		}
	}
	return
}

func (d *Dao)CountForUserAddress(userId int64) (err error) {
	var count int
	err = d.db.Model(&model.UserAddress{}).Where("user_id = ? and status = 0", userId).Count(&count).Error
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", userId, err)
		return
	}
	return
}

func (d *Dao)ListForUserAddress(userId int64) (userAddresses []model.UserAddress, err error) {
	err = d.db.Where("user_id = ? and status = 0", userId).Find(&userAddresses).Error
	if err != nil {
		log.Error("d.db.Where(%v) err(%v)", userId, err)
		return
	}
	return
}
