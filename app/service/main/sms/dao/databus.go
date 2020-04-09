package dao

import (
	v1 "angrymiao-go/app/service/main/sms/api"
	"angrymiao-go/punk/log"
	"context"
	"time"
)
const _retry = 3

// PubSingle pub single sms to databus
func (d *Dao) PubSingle(ctx context.Context, l *v1.SendReq) (err error) {
	for i := 0; i < _retry; i++ {
		if err = d.databus.Send(ctx, l.Mobile, l); err == nil {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if err != nil {
		log.Error("PubSingle(%+v) error(%v)", l, err)
		return
	}
	log.Info("PubSingle(%+v) success.", l)
	return
}