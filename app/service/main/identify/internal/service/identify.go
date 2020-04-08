package service

import (
	v1 "angrymiao-go/app/service/main/identify/api"
	"angrymiao-go/app/service/main/identify/internal/model"
	"angrymiao-go/punk/ecode"
	"angrymiao-go/punk/net/metadata"
	"context"
)

const (
	// cache -101
	_noLoginMid         = 0
	_CookieSessionField = "SESSDATA"
	_CookieBuvidField   = "buvid3"
)

var (
	_noLoginIdentify = &model.IdentifyInfo{
		Mid:     _noLoginMid,
		Expires: 86400,
	}
)

// GetTokenInfo verify user info by accesskey.
func (s *Service) GetTokenInfo(c context.Context, token *v1.GetTokenInfoReq) (res *model.IdentifyInfo, err error) {
	var cache = true
	if res, err = s.d.AccessCache(c, token.Token); err != nil {
		cache = false
	}
	if res != nil {
		if res.Mid == _noLoginMid {
			err = ecode.NoLogin
			return
		}
		return
	}
	if res, err = s.d.AccessToken(c, token.Token); err != nil {
		if err != ecode.NoLogin && err != ecode.AccessKeyErr {
			return
		}
		// no login and need cache 30s
		res = _noLoginIdentify
	}
	if cache && res != nil {
		s.cache.Save(func() {
			s.d.SetAccessCache(context.Background(), token.Token, res)
		})
	}
	if res.Mid == _noLoginMid {
		err = ecode.NoLogin
		return
	}
	return
}
