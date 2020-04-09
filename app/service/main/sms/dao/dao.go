package dao

import (
	"angrymiao-go/app/service/main/sms/conf"
	"angrymiao-go/punk/queue/databus"
)

// Dao .
type Dao struct {
	c       *conf.Config
	databus *databus.Databus
}

// New new a instance
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		databus: databus.New(c.Databus),
	}
	return
}

// Close close connection of db , mc.
func (d *Dao) Close() {
	d.databus.Close()
}
