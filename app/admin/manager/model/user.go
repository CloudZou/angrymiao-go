package model

import "angrymiao-go/punk/model"

// const rbac const
const (
	// auth_item type
	TypePointer  = 1
	TypeCategory = 2
	TypeRole     = 3
	TypeGroup    = 4

	// Admin super admin
	Admin = 1
	// user state
	UserStateOn = 0
	UserDepOn   = 1
)

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