package dao

import (
	"angrymiao-go/app/job/sms/conf"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/queue/databus"
	"angrymiao-go/punk/stat/prom"
)

var (
	errorsCount = prom.BusinessErrCount
	infosCount  = prom.BusinessInfoCount
)

// Dao .
type Dao struct {
	c       *conf.Config
	databus *databus.Databus
	httpClient *bm.Client
}

// New new a instance
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
		databus: databus.New(c.Databus),
		httpClient: bm.NewClient(c.HTTPClient),
	}
	return
}

// PromError prometheus error count.
func PromError(name string) {
	errorsCount.Incr(name)
}

// PromInfo prometheus info count.
func PromInfo(name string) {
	infosCount.Incr(name)
}


// Close close connection of db , mc.
func (d *Dao) Close() {
	d.databus.Close()
}
