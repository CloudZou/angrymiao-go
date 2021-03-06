package http

import (
	"angrymiao-go/app/admin/passport/conf"
	"angrymiao-go/app/admin/passport/model"
	"angrymiao-go/app/admin/passport/service"
	"angrymiao-go/punk/ecode"
	"angrymiao-go/punk/net/http/blademaster/middleware/permit"
	"encoding/hex"
	"math/rand"
	"net/http"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"net/url"
	"time"
)

var (
	svc *service.Service
)

const (
	_sessIDKey             = "_AJSESSIONID"
	_defaultDomain         = ".angrymiao.com"
	_defaultCookieLifeTime = 2592000
	_sessionIDLength 	   = 32
)

// New new a bm server.
func New(c *conf.Config, s *service.Service) (engine *bm.Engine, err error) {
	engine = bm.DefaultServer(c.Server)
	svc = s
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/x/admin/passport")
	{
		g.POST("/signIn", signIn)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func signIn(c *bm.Context) {
	var (
		err error
	)
	var signInForm = &model.SignInForm{}
	if err = c.Bind(signInForm); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	res := newSession()
	res.Values["username"] = signInForm.Username
	err = svc.SignIn(c, res.Sid, signInForm)
	if err != nil {
		c.JSON("", ecode.AuthFailed)
		return
	}
	setHTTPCookie(c, _sessIDKey, res.Sid)
	c.JSON("", nil)
}

func newSession() (res *permit.Session) {
	b := make([]byte, _sessionIDLength)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return nil
	}
	res = &permit.Session{
		Sid:    hex.EncodeToString(b),
		Values: make(map[string]interface{}),
	}
	return
}

func setHTTPCookie(ctx *bm.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		Path:     "/",
		HttpOnly: true,
		Domain:   _defaultDomain,
	}
	cookie.MaxAge = _defaultCookieLifeTime
	cookie.Expires = time.Now().Add(time.Duration(_defaultCookieLifeTime) * time.Second)
	http.SetCookie(ctx.Writer, cookie)
}