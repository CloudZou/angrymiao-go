package ecode

var (
	//sms
	SmsMobilePatternErr       = New(66031) // 手机号码格式不正确
	SmsMobileFrequencyErr	  = New(66032)  //发送验证码频率过快
)