package model

import "angrymiao-go/punk/model"

// const rbac const
const (
)

type User struct {
	Username     string    `json:"username"`
	model.BaseModel
}

// Account auth user account
type Account struct {
	Username string `json:"username"`
	model.BaseModel
}

// Auth .
type Auth struct {
	UID        int64    `json:"uid"`
	Username   string   `json:"username"`
	Perms      []string `json:"perms"`
	Admin      bool     `json:"admin"`
}
// Permissions .
type Permissions struct {
	UID   int64      `json:"uid"`
	Perms []string   `json:"perms"`
	Admin bool       `json:"admin"`
}