package dao

import (
	"context"

	"angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/rate/limit/bench/stress/conf"
)

// Dao dao
type Dao struct {
	c     *conf.Config
	redis *redis.Pool
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:     c,
		redis: redis.NewPool(c.Redis),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return d.pingMC(c)
}

// pingMc ping
func (d *Dao) pingMC(c context.Context) (err error) {
	return
}
