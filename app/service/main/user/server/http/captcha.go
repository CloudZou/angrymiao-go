package http

import (
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/punk/ecode"
	bm "angrymiao-go/punk/net/http/blademaster"
)

func sendCode(c *bm.Context) {
	phone := c.Request.Form.Get("phone")
	if err := srv.ValidateAndSend(c, phone, model.SIGN_UP); err != nil {
		c.JSON(nil, ecode.SignInCodeSendFailed)
		return
	}
	c.JSON(nil, nil)
	return
}
