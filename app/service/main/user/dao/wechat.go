package dao

import (
	"angrymiao-go/punk/log"
	"github.com/silenceper/wechat/oauth"
)

func (d *Dao)OAuth2GetUserInfo(code string) (userInfo oauth.UserInfo, err error) {
	oauth := d.wechat.GetOauth()
	resToken, err := oauth.GetUserAccessToken(code)
	if err != nil {
		log.Error("oauth().GetUserAccessToken(%v) err(%v)", code, err)
		return
	}
	userInfo, err = oauth.GetUserInfo(resToken.AccessToken, resToken.OpenID)
	if err != nil {
		log.Error("oauth.GetUserInfo(%v, %v) err(%v)", resToken.AccessToken, resToken.OpenID)
		return
	}
	return
}
