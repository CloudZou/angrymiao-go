package model

import "angrymiao-go/punk/model"

type Account struct {
	Username string		`json:"username" gorm:"user_name"`
	HashedPassword string 		`json:"hashedPassword" gorm:"hashed_password"`

	model.BaseModel
}

func (a *Account) TableName() string {
	return "account"
}