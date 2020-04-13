package http

import (
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/app/service/main/user/model"
	"angrymiao-go/app/service/main/user/service"
	"angrymiao-go/punk/ecode"
	"angrymiao-go/punk/jwt"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"fmt"
)

const (
	AUTHORIZATION_HEADER = "Authorization"
)
var srv *service.Service

// New new a bm server.
// New new a bm server.
func New(c *conf.Config, s *service.Service) (engine *bm.Engine, err error) {
	srv = s
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
	g := e.Group("/identify")
	{

	}
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
	c.Writer.Header().Add(AUTHORIZATION_HEADER, fmt.Sprintf("Bearer %s", accessToken))
	userViewResponse := user.ToUserViewResponse()
	userViewResponse.Token = accessToken
	userMap := make(map[string]interface{})
	userMap["user"] = userViewResponse
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
	c.Writer.Header().Add(AUTHORIZATION_HEADER, fmt.Sprintf("Bearer %s", accessToken))
	userViewResponse := user.ToUserViewResponse()
	userViewResponse.Token = accessToken
	userMap := make(map[string]interface{})
	userMap["user"] = userViewResponse
	c.JSON(userMap, nil)
	return
}