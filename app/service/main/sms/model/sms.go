package model

type AliyunSmsRequest struct {
	Phone string `json:"phone"`
	SignName string `json:"signName"`
	TemplateCode string `json:"templateCode"`
	AccessKeyId string `json::"accessKeyId"`
	Action string `json:"action"`
	OutId string `json:"outId"`
	SmsUpExtendCode string `json::smsUpExtendCode`
	TemplateParam string `json::"templateParam"`
}