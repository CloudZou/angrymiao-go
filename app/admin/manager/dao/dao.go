package dao

import (
	"angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/log"
	"context"

	"angrymiao-go/app/admin/manager/conf"
	"angrymiao-go/punk/database/orm"
	bm "angrymiao-go/punk/net/http/blademaster"

	"github.com/jinzhu/gorm"
)

// Dao .
type Dao struct {
	db         *gorm.DB
	redis         *redis.Pool
	httpClient *bm.Client
	dsbClient  *bm.Client
}

// New new a instance
func New(c *conf.Config) (d *Dao) {
	db, _, err := orm.NewMySQL(c.ORM)
	if err != nil {
		log.Error("orm.NewMySQL(%v) error(%v)",c.ORM, err)
		return
	}
	d = &Dao{
		db:         db,
		redis:      redis.NewPool(c.Redis),
		httpClient: bm.NewClient(c.HTTPClient),
		dsbClient:  bm.NewClient(c.DsbClient),
	}
	d.initORM()
	return
}

func (d *Dao) initORM() {
	d.db.LogMode(true)
}

// DB .
func (d *Dao) DB() *gorm.DB {
	return d.db
}

// Ping check connection of db , mc.
func (d *Dao) Ping(c context.Context) (err error) {
	if d.db != nil {
		err = d.db.DB().PingContext(c)
	}
	return
}

// Close close connection of db , mc.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
