package dao

import (
	"angrymiao-go/app/admin/passport/conf"
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
}

func NewDao(c *conf.Config) (d *Dao, cf func(), err error) {
	db, cf, err := orm.NewMySQL(c.DB)
	if err != nil {
		log.Error("NewDB(%v) err(%v)", c, err)
	}

	d = &Dao{
		db: db,
		cache: fanout.New("cache"),
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
