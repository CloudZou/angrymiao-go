package permit

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"sync"
	"time"

	"angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
)

// Session http session.
type Session struct {
	Sid string

	lock   sync.RWMutex
	Values map[string]interface{}
}

// SessionConfig config of Session.
type SessionConfig struct {
	SessionIDLength int
	CookieLifeTime  int
	CookieName      string
	Domain          string

	RedisConfig *redis.Config
}

// SessionManager .
type SessionManager struct {
	rp *redis.Pool // Session cache
	c  *SessionConfig
}

// newSessionManager .
func newSessionManager(c *SessionConfig) (s *SessionManager) {
	s = &SessionManager{
		rp: redis.NewPool(c.RedisConfig),
		c:  c,
	}
	return
}

// SessionStart start session.
func (s *SessionManager) SessionStart(ctx *bm.Context) (si *Session) {
	// check manager Session id, if err or no exist need new one.
	if si, _ = s.cache(ctx); si == nil {
		si = s.newSession(ctx)
	}
	return
}

// SessionRelease flush session into store.
func (s *SessionManager) SessionRelease(ctx *bm.Context, sv *Session) {
	// set http cookie
	s.setHTTPCookie(ctx, s.c.CookieName, sv.Sid)
	// set mc
	conn := s.rp.Get(ctx)
	defer conn.Close()
	key := sv.Sid
	data, err := json.Marshal(sv)
	if err != nil {
		log.Error("json.Marshal(%v) err(%v)", sv, err)
	}
	conn.Send("SET", key, data)
	conn.Send("EXPIRE", key, int32(s.c.CookieLifeTime))
	if _, err := conn.Do("EXEC"); err != nil {
		log.Error("SessionManager set error(%s,%v)", key, err)
	}
}

// SessionDestroy destroy session.
func (s *SessionManager) SessionDestroy(ctx *bm.Context, sv *Session) {
	conn := s.rp.Get(ctx)
	defer conn.Close()
	if _, err := conn.Do("DEL", sv.Sid); err != nil {
		log.Error("SessionManager delete error(%s,%v)", sv.Sid, err)
	}
}

func (s *SessionManager) cache(ctx *bm.Context) (res *Session, err error) {
	ck, err := ctx.Request.Cookie(s.c.CookieName)
	if err != nil || ck == nil {
		return
	}
	sid := ck.Value
	// get from cache
	conn := s.rp.Get(ctx)
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("GET", sid))
	if err != nil {
		log.Error("conn.Do(GET, %s) error(%v)", sid, err)
		return
	}
	res = &Session{}
	if err = json.Unmarshal(reply, res); err != nil {
		log.Error("json.Unmarshal(%v,%v) error(%v)", string(reply), res, err)
	}
	return
}

func (s *SessionManager) newSession(ctx context.Context) (res *Session) {
	b := make([]byte, s.c.SessionIDLength)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return nil
	}
	res = &Session{
		Sid:    hex.EncodeToString(b),
		Values: make(map[string]interface{}),
	}
	return
}

func (s *SessionManager) setHTTPCookie(ctx *bm.Context, name, value string) {
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

// Get get value by key.
func (s *Session) Get(key string) (value interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	value = s.Values[key]
	return
}

// Set set value into session.
func (s *Session) Set(key string, value interface{}) (err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Values[key] = value
	return
}
