package dao

import (
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/database/orm"
	"angrymiao-go/punk/log"
	"angrymiao-go/punk/sync/pipeline/fanout"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/silenceper/wechat"
)

// Dao Dao.
type Dao struct {
	db    *gorm.DB
	redis *redis.Pool
	cache *fanout.Fanout
	wechat *wechat.Wechat
	QQ     *QQConnector
}

func New(c *conf.Config) (d *Dao, cf func(), err error) {
	db, _, err := orm.NewMySQL(c.Db)
	if err != nil {
		log.Error("orm.NewMySQL(%v) error(%v)",c.Db, err)
		return
	}
	d = &Dao{
		db:         db,
		redis:      redis.NewPool(c.Redis),
		wechat: wechat.NewWechat(c.Wechat),
		QQ: &QQConnector{Config:c.QQ},
	}
	d.initORM()
	return
}

func (d *Dao) initORM() {
	d.db.LogMode(true)
}

// Close close the resource.
func (d *Dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}
