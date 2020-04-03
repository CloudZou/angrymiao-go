package model

import "angrymiao-go/punk/model"

type Account struct {
	Username string		`json:"username" orm:""`
	HashedPassword string 		`json:"hashedPassword"`

	model.BaseModel
}

func (a *Account) TableName() string {
	return "account"
}

type AccountSession struct {
	AccountID string 	`json:"accountId" gorm:"account_id"`
	SessionID string 	`json:"sessionId" gorm:"session_id"`

	model.BaseModel
}

func (a *AccountSession) TableName() string {
	return "account_session"
}