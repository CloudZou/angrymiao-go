package model

type PhoneLoginReq struct {
	Phone string `json::"phone"`
	Code string `json:"code"`
}
