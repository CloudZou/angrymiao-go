package dao

import (
	"angrymiao-go/punk/cache/redis"
	"context"
	"encoding/json"

	"angrymiao-go/punk/log"
	"angrymiao-go/punk/net/http/blademaster/middleware/permit"
)

// Session .
func (d *Dao) Session(ctx context.Context, sid string) (res *permit.Session, err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	data, err := redis.Bytes(conn.Do("GET", sid))
	res = &permit.Session{}
	if err = json.Unmarshal(data, res); err != nil {
		log.Error("json.Unmarshal(%v,%v) errror(%v)", string(data), res, err)
	}
	return
}

// SetSession .
func (d *Dao) SetSession(ctx context.Context, p *permit.Session) (err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	key := p.Sid
	data, err := json.Marshal(p)
	if err != nil {
		log.Error("json.Marshal(%v) err(%v)", p, err)
	}
	conn.Send("SET", key, data)
	conn.Send("EXPIRE", key, int32(_sessionLife))
	if _, err := conn.Do("EXEC"); err != nil {
		log.Error("conn.Send set error(%s,%v)", key, err)
	}
	return
}
