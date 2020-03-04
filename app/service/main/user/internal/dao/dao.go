package dao

import (
	"context"
	"github.com/CloudZou/punk/pkg/gredis"
	"github.com/CloudZou/punk/pkg/sync/pipeline/fanout"
	"github.com/jinzhu/gorm"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedisClient)

// Dao Dao.
type Dao struct {
	db    *gorm.DB
	redis *gredis.RedisClient
	cache *fanout.Fanout
}

func New(r *gredis.RedisClient, db *gorm.DB) (d *Dao, cf func(), err error) {
	d = &Dao{
		db:    db,
		redis: r,
		cache: fanout.New("cache"),
	}
	cf = d.Close
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
