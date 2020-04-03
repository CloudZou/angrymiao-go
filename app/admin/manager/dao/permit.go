package dao

import (
	"angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/log"
	"angrymiao-go/punk/net/http/blademaster/middleware/permit"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
)

const (
	_sessionlen   = 32
	_sessionLife  = 2592000
)

// VerifyDsb .
func (d *Dao) VerifyDsb(ctx context.Context, sid string) (res string, err error) {
	var dsbRes struct {
		Code     int    `json:"code"`
		UserName string `json:"username"`
	}
	conn := d.redis.Get(ctx)
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("GET", sid))
	if err != nil {
		log.Error("conn.Do(GET, %s) error(%v)", sid, err)
		return
	}
	if err = json.Unmarshal(reply, &dsbRes); err != nil {
		log.Error("json.Unmarshal(%v,%v) error(%v)", string(reply), res, err)
	}
	//从redis取数据判断
	res = dsbRes.UserName
	return
}

// NewSession .
func (d *Dao) NewSession(ctx context.Context) (res *permit.Session) {
	b := make([]byte, _sessionlen)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return
	}
	res = &permit.Session{
		Sid:    hex.EncodeToString(b),
		Values: make(map[string]interface{}),
	}
	return
}
