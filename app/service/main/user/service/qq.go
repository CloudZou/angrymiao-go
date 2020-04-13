package service

import (
	"angrymiao-go/punk/log"
	"fmt"
	"io/ioutil"
	xhttp "net/http"
	"net/url"
)

const (
	QQ_GRAPH_HOST = "https://graph.qq.com"
	QQ_GRPAH_URL = "%s/oauth2.0/token?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s"
)

type QQConfig struct {
	AppId 	string
	AppKey  string

}

type QQConnector struct {
	Config *QQConfig
}

func (c *QQConnector)getQQAccessToken(code string, redirectUrl string) {
	fullUrl := fmt.Sprintf(QQ_GRPAH_URL, QQ_GRAPH_HOST, c.Config.AppId, c.Config.AppKey, code, url.QueryEscape(redirectUrl))
	resp, err := xhttp.Get(fullUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Error("xhttp.Get(%v) err(%v)", fullUrl, err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
}