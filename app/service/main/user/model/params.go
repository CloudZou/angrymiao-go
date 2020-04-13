package model

import "time"

type PhoneLoginReq struct {
	Phone string `json::"phone"`
	Code string `json:"code"`
}

type UserViewResponse struct {
	ID int64 	`json:"id"`
	Nickname string 	`json:"nickname"`
	Phone string 	`json:"phone"`
	AvatarUrl string 	`json:"avatar_url"`
	Balance string 	`json:"balance"`
	Gender int	`json:"gender"`
	DateJoined time.Time	`json:"date_joined"`
	AuthType AuthType	`json:"auth_type"`
	Token string 	`json:"token"`
	RefreshToken string `json:refreshToken`
}

type WechatInfoReq struct {
	Code string 	`json:"code"`
	Scope string 	`json:"scope"`
	RedirectUrl string `json:"redirect_url"`
	Scene string `json:"scene"`
}