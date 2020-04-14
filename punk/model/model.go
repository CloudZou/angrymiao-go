package model

import "time"

type BaseModel struct {
	ID         int64 `gorm:"column:id;primary_key" json:"id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	Status int `gorm:"status" json:"status"`
}

