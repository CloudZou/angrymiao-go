package dao

import (
	"angrymiao-go/app/admin/passport/model"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"encoding/json"
	"time"
)
const (
	_defaultCookieLifeTime = 2592000
)

func (d *Dao) Auth(ctx *bm.Context, key string, signInForm *model.SignInForm) (err error) {
	var count int
	err = d.db.Where("account = ? and hashed_password= ?", signInForm.Username, signInForm.Password).Count(&count).Error
	if err != nil {
		log.Error("d.db.Where(%v,%v)", signInForm.Username, signInForm.Password)
		return
	}
	if count == 1 {
		err = d.addCache(ctx, key, signInForm)
		if err != nil {
			log.Error("d.addCache(ctx, %v, %v) err(%v)", key, signInForm, err)
			return
		}
		return
	}
	return
}

func (d *Dao)addCache(ctx *bm.Context, key string, signInForm *model.SignInForm) (err error) {
	var cacheForm struct {
		Username string 	`json:"username"`
	}
	cacheForm.Username = signInForm.Username
	data, err :=json.Marshal(cacheForm)
	if err != nil {
		log.Error("json.Marshal(%v) err(%v)", cacheForm, err)
		return
	}
	conn := d.redis.Get(ctx)
	defer conn.Close()
	_, err = conn.Do("SET", key, data)
	conn.Send("EXPIRE", key, time.Duration(_defaultCookieLifeTime) * time.Second)
	if err != nil {
		log.Error("conn.Do(SET, %v, %v)", key, string(data))
		return
	}
	return
}