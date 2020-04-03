package model

import "time"

type BaseModel struct {
	ID         int `orm:"primary_key" json:"id"`
	CreateTime time.Time `orm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `orm:"column:update_time" json:"update_time"`
	Status int `orm:"status" json:"status"`
}

