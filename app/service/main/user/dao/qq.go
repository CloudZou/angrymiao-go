package dao

import (
	"angrymiao-go/punk/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	. "angrymiao-go/app/service/main/user/model"
	. "angrymiao-go/app/service/main/user/conf"
	xhttp "net/http"
)


const (
	QQ_GRAPH_HOST = "https://graph.qq.com"
	QQ_GRPAH_ACCESS_TOKEN_URL = "%s/oauth2.0/token?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s"
	QQ_GRPAH_OPEN_ID_URL = "%s/oauth2.0/me?access_token=%s&unionid=%d"
	QQ_GRAPH_USER_INFO_URL = "%s/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s"
	APPLY_UNION_ID = 1
)

type QQConnector struct {
	Config *QQConfig
}

func (c *QQConnector)GetQQAccessToken(code string, redirectUrl string) (qqTokenResponse QQTokenResponse, err error){
	fullUrl := fmt.Sprintf(QQ_GRPAH_ACCESS_TOKEN_URL, QQ_GRAPH_HOST, c.Config.AppId, c.Config.AppKey, code, url.QueryEscape(redirectUrl))
	resp, err := xhttp.Get(fullUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Error("xhttp.Get(%v) err(%v)", fullUrl, err)
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	jsonStr := getJsonStr(data)
	if err = json.Unmarshal([]byte(jsonStr), &qqTokenResponse); err != nil {
		log.Error("json.Unmarshal(%v, %v) err(%v)", jsonStr, qqTokenResponse)
		return
	}
	return
}

func (c *QQConnector)GetQQOpenId(accessToken string) (qqOpenIDResponse QQOpenIDResponse, err error) {
	fullUrl := fmt.Sprintf(QQ_GRPAH_OPEN_ID_URL, QQ_GRAPH_HOST, accessToken, APPLY_UNION_ID)
	resp, err := xhttp.Get(fullUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Error("xhttp.Get(%v) err(%v)", fullUrl, err)
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	jsonStr := getJsonStr(data)
	if err = json.Unmarshal([]byte(jsonStr), &qqOpenIDResponse); err != nil {
		log.Error("json.Unmarshal(%v, %v) err(%v)", jsonStr, qqOpenIDResponse)
		return
	}
	return
}

func (c *QQConnector) GetQQUserInfo(accessToken string, openID string)(qqUserInfo QQUserInfoResponse, err error) {
	fullUrl := fmt.Sprintf(QQ_GRAPH_USER_INFO_URL, QQ_GRAPH_HOST, accessToken, c.Config.AppId, openID)
	resp, err := xhttp.Get(fullUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Error("xhttp.Get(%v) err(%v)", fullUrl, err)
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	jsonStr := getJsonStr(data)
	if err = json.Unmarshal([]byte(jsonStr), &qqUserInfo); err != nil {
		log.Error("json.Unmarshal(%v, %v) err(%v)", jsonStr, qqUserInfo)
		return
	}
	return
}

func getJsonStr(data []byte)(jsonStr string) {
	body := string(data)
	startIndex := strings.Index(body, "{")
	endIndex := strings.Index(body, "}")
	jsonStr = substring(body, startIndex, endIndex)
	return
}


func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)
	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	var substring = ""
	for i := start; i < length; i++ {
		substring += string(r[i])
	}

	return substring
}