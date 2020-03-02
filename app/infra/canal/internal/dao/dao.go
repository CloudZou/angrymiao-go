package dao

import (
	"angrymiao-go/app/infra/canal/conf"
	"github.com/jinzhu/gorm"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB)

// Dao Dao.
type Dao struct {
	db *gorm.DB
	c  *conf.Config
}

func New(db *gorm.DB) (d *Dao, err error) {
	d = &Dao{
		c:  conf.Conf,
		db: db,
	}
	return
}
