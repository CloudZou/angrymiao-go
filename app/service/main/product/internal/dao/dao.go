package dao

import (
	"context"
	"angrymiao-go/punk/sync/pipeline/fanout"
	"github.com/jinzhu/gorm"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB)

// Dao Dao.
type Dao struct {
	db    *gorm.DB
	cache *fanout.Fanout
}

func New(db *gorm.DB) (d *Dao, cf func(), err error) {
	d = &Dao{
		db:    db,
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
