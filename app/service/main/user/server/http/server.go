package http

import (
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/app/service/main/user/service"
	"angrymiao-go/punk/ecode"
	"angrymiao-go/punk/jwt"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	authService "angrymiao-go/punk/net/http/blademaster/middleware/auth"
	"fmt"
)

const (
	AUTHORIZATION_HEADER = "Authorization"
	AUTHORIZATION_BEARER = "Bearer"
)
var srv *service.Service
var auth *authService.Auth
// New new a bm server.
func New(c *conf.Config, s *service.Service) (engine *bm.Engine, err error) {
	srv = s
	auth = authService.New()
	engine = bm.DefaultServer(c.BM)
	initRouter(engine)
	err = engine.Start()
	if err := engine.Start(); err != nil {
		log.Error("engine.Start error(%v)", err)
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	g := e.Group("/user")
	{
		g.POST("/phoneLogin", phoneLogin)
		g.POST("/wechatLogin", wechatLogin)
		g.POST("/qqLogin", qqLogin)

		g.POST("/updateWxOpenId", auth.UserMobile, updateWxOpenId)
	}
}

func updateWxOpenId(c *bm.Context) {
	var err error
	user, err := srv.GetCurrentUser(c)
	if err != nil {
		c.JSON(nil, ecode.NoLogin)
		return
	}
	p := new(model.UpdateWxOpenIdReq)
	if err = c.Bind(p); err != nil {
		return
	}
	err = srv.UpdateWxOpenIdById(*user, p.WxOpenId)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(nil, nil)
	return
}

func qqLogin(c *bm.Context) {
	p := new(model.WechatInfoReq)
	if err := c.Bind(p); err != nil {
		return
	}
	user, err := srv.QQLogin(c, p)
	if err != nil {
		log.Error("svr.QQLogin(c,%v) err(%v)", p, err)
		c.JSON(nil, ecode.AuthCodeCheckFailed)
		return
	}
	accessToken, err := jwt.GenerateToken(user.Phone, user.ID)
	if err != nil {
		log.Error("jwt.GenerateToken(%v,%v) err(%v)", user.Phone, user.ID, err)
		c.JSON(nil, ecode.AuthCodeCheckFailed)
		return
	}
	c.Writer.Header().Add(AUTHORIZATION_HEADER, fmt.Sprintf("%s %s",AUTHORIZATION_BEARER, accessToken))
	userMap := renderUserProfile(user, accessToken)
	c.JSON(userMap, nil)
	return
}

func phoneLogin(c *bm.Context) {
	p := new(model.PhoneLoginReq)
	if err := c.Bind(p); err != nil {
		return
	}
	user, err := srv.PhoneLogin(c, p)
	if err != nil {
		log.Error("svr.PhoneLogin(c,%v) err(%v)", p, err)
		c.JSON(nil, ecode.AuthCodeCheckFailed)
		return
	}
	accessToken, err := jwt.GenerateToken(user.Phone, user.ID)
	if err != nil {
		log.Error("jwt.GenerateToken(%v,%v) err(%v)", user.Phone, user.ID, err)
		c.JSON(nil, ecode.AuthCodeCheckFailed)
		return
	}
	c.Writer.Header().Add(AUTHORIZATION_HEADER, fmt.Sprintf("%s %s", AUTHORIZATION_BEARER, accessToken))
	userMap := renderUserProfile(user, accessToken)
	c.JSON(userMap, nil)
	return
}

func wechatLogin(c *bm.Context) {
	p := new(model.WechatInfoReq)
	if err := c.Bind(p); err != nil {
		return
	}
	user, err := srv.WechatLogin(c, p)
	if err != nil {
		log.Error("svr.WechatLogin(c,%v) err(%v)", p, err)
		c.JSON(nil, ecode.AuthCodeCheckFailed)
		return
	}
	accessToken, err := jwt.GenerateToken(user.Phone, user.ID)
	if err != nil {
		log.Error("jwt.GenerateToken(%v,%v) err(%v)", user.Phone, user.ID, err)
		c.JSON(nil, ecode.AuthCodeCheckFailed)
		return
	}
	c.Writer.Header().Add(AUTHORIZATION_HEADER, fmt.Sprintf("%s %s", AUTHORIZATION_BEARER, accessToken))
	userMap := renderUserProfile(user, accessToken)
	c.JSON(userMap, nil)
	return
}

func renderUserProfile(user *model.User, accessToken string) (userMap map[string]interface{}) {
	userViewResponse := user.ToUserViewResponse()
	userViewResponse.Token = accessToken
	userMap = make(map[string]interface{})
	userMap["user"] = userViewResponse
	return
}