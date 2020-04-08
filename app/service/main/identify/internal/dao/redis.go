package dao

import (
	"angrymiao-go/app/service/main/identify/internal/model"
	"context"
	"fmt"
)

var (
	loginCacheValue = []byte("1")
)

// SetAccessCache .
func (d *Dao) SetAccessCache(c context.Context, key string, res *model.IdentifyInfo) {
	conn := d.redis.Get(c)
	defer conn.Close()
	key = cacheKey(key)

}

// AccessCache .
func (d *Dao) AccessCache(c context.Context, key string) (res *model.IdentifyInfo, err error) {
	conn := d.redis.Get(c)
	defer conn.Close()
	key = cacheKey(key)
	//r, err := conn.Get(key)

	//res = &model.IdentifyInfo{}
	//if err = conn.Scan(r, res); err != nil {
	//	PromError("mc:json解析失败")
	//	log.Error("conn.Scan(%v) error(%v)", string(r.Value), err)
	//	return
	//}
	cachedCount.Incr("access_cache")
	return
}

// DelCache delete access cache.
func (d *Dao) DelCache(c context.Context, key string) (err error) {
	conn := d.redis.Get(c)
	defer conn.Close()
	key = cacheKey(key)
	return
}

func cacheKey(key string) string {
	return fmt.Sprintf("i_%s", key)
}

func loginCacheKey(mid int64, ip string) string {
	return fmt.Sprintf("l%d%s", mid, ip)
}
