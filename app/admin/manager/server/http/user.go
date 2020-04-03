package http

import (
	"angrymiao-go/punk/ecode"
	bm "angrymiao-go/punk/net/http/blademaster"
	"net/http"
)

const _logout = "http://auth-mng.bilibili.co/logout?caller=manager-admin"

func authUser(c *bm.Context) {
	var (
		username string
	)
	if un, ok := c.Get("username"); ok {
		username = un.(string)
	} else {
		c.JSON(nil, ecode.Unauthorized)
		return
	}
	c.JSON(mngSvc.Auth(c, username))
}

func logout(c *bm.Context) {
	// purge mid cache
	c.Redirect(http.StatusFound, _logout)
}

func permissions(c *bm.Context) {
	var username string
	if username = c.Request.Form.Get("username"); username == "" {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(mngSvc.Permissions(c, username))
}
