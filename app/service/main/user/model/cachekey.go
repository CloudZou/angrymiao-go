package model

import "fmt"

const (
	PREFIX_ALL string = "all_"
)

type SmsType string
const (
	SIGN_UP SmsType = "SIGN_UP"
	SIGN_IN SmsType = "SIGN_IN"
	BIND_MOBILE SmsType = "BIND_MOBILE"
)


func KeyOfMobileCaptcha(mobile string, smsType SmsType) string {
	return fmt.Sprintf("%smobileCaptcha#mobile=%s#smsType=%s", PREFIX_ALL, mobile, smsType)
}

func KeyOfMobileGetCaptchaFrequency(mobile string, smsType SmsType) string {
	return fmt.Sprintf("%mobileGetCaptchaFrequency#mobile=%s#smsType=%s", PREFIX_ALL, mobile, smsType)
}

func KeyOfMobileCaptchaTryTimes(mobile string, smsType SmsType) string  {
	return fmt.Sprintf("%mobileCaptchaTryTimes#mobile=%s#smsType=%s", PREFIX_ALL, mobile, smsType)
}