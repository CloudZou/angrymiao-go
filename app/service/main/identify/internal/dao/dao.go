package dao

import (
	"angrymiao-go/app/service/main/identify/conf"
	"angrymiao-go/punk/cache/redis"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/stat/prom"
	"context"
)


const (
	//h5 统一登陆验证入口
	_tokenURI  = "/intranet/auth/tokenInfo"
	_cookieURI = "/intranet/auth/cookieInfo"
)

var (
	errorsCount = prom.BusinessErrCount
	cachedCount = prom.CacheHit
	missedCount = prom.CacheMiss
)

// Dao Dao.
type Dao struct {
	c *conf.Config
	redis *redis.Pool

	tokenURI  string
	cookieURI string
	client    *bm.Client
}


func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:         c,
		tokenURI:  c.Identify.AuthHost + _tokenURI,
		cookieURI: c.Identify.AuthHost + _cookieURI,
		client:    bm.NewClient(c.HTTPClient),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {

}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}
