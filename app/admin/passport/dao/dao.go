package dao

import (
	"angrymiao-go/app/admin/passport/conf"
	redis "angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/database/orm"
	"angrymiao-go/punk/log"
	"angrymiao-go/punk/sync/pipeline/fanout"
	"context"
	"github.com/jinzhu/gorm"
)

// Dao Dao.
type Dao struct {
	db          *gorm.DB
	cache *fanout.Fanout
	redis *redis.Pool
}

func NewDao(c *conf.Config) (d *Dao, cf func(), err error) {
	db, cf, err := orm.NewMySQL(c.DB)
	if err != nil {
		log.Error("NewDB(%v) err(%v)", c, err)
	}

	r := redis.NewPool(c.Redis)

	d = &Dao{
		db: db,
		cache: fanout.New("cache"),
		redis: r,
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}
