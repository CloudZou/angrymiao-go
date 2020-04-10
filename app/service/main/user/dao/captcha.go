package dao

import (
	"angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/ecode"
	"angrymiao-go/punk/log"
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	_frequencyExpire = 3 * 60
	_captchaExpire = 5 * 60
	_length = 6
)

func (d *Dao)ValidateFrequency(c context.Context, key string) (err error) {
	conn := d.redis.Get(c)
	defer conn.Close()
	frequency, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		log.Error("conn.Do(GET,%v) err(%v)", key, err)
		return
	}
	if frequency > 3 {
		err = ecode.SmsMobileFrequencyErr
		return
	}
	conn.Do("INCR", key)
	return
}

func (d *Dao)InitFrequencyCache(c context.Context, key string) (err error) {
	conn := d.redis.Get(c)
	defer conn.Close()
	//TODO 需要跟下面的方法一起考虑重构
	if err = conn.Send("SET", key, 1); err != nil {
		log.Error("redis: conn.Send(SET, %s, %s) error(%v)", key, 1, err)
		return
	}
	if err = conn.Send("EXPIRE", key, time.Duration(_frequencyExpire) * time.Second); err != nil {
		log.Error("redis: conn.Send(EXPIRE key(%s) expire(%d)) error(%v)", key, _frequencyExpire, err)
		return
	}
	if err = conn.Flush(); err != nil {
		log.Error("redis: conn.Flush error(%v)", err)
		return
	}
	return
}

func (d *Dao)AddCaptchaCache(c context.Context, key string, captcha string) (err error) {
	conn := d.redis.Get(c)
	defer conn.Close()
	if err = conn.Send("SET", key, captcha); err != nil {
		log.Error("redis: conn.Send(SET, %s, %s) error(%v)", key, captcha, err)
		return
	}
	if err = conn.Send("EXPIRE", key, time.Duration(_captchaExpire) * time.Second); err != nil {
		log.Error("redis: conn.Send(EXPIRE key(%s) expire(%d)) error(%v)", key, _frequencyExpire, err)
		return
	}
	if err = conn.Flush(); err != nil {
		log.Error("redis: conn.Flush error(%v)", err)
		return
	}
	return
}

func (d *Dao)GetCaptcha(c context.Context, key string)(captcha string, err error) {
	conn := d.redis.Get(c)
	defer conn.Close()
	captcha, err = redis.String(conn.Do("GET", key))
	if err != nil {
		log.Error("redis: conn.Do(GET, %s) error(%v)", key, err)
		return
	}
	return
}


func (d *Dao)GenerateRandomNumericString() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}