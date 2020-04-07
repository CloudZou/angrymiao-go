package permit_test

import (
	"fmt"
	"time"

	"angrymiao-go/punk/cache/memcache"
	"angrymiao-go/punk/container/pool"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/http/blademaster/middleware/permit"
	"angrymiao-go/punk/net/metadata"
	"angrymiao-go/punk/net/netutil/breaker"
	xtime "angrymiao-go/punk/time"
)

// This example create a permit middleware instance and attach to several path,
// it will validate request by specified policy and put extra information into context. e.g., `uid`.
// It provides additional handler functions to provide the identification for your business handler.
func Example() {
	a := permit.New(&permit.Config{
		DsHTTPClient: &bm.ClientConfig{
			App: &bm.App{
				Key:    "manager-go",
				Secret: "949bbb2dd3178252638c2407578bc7ad",
			},
			Dial:      xtime.Duration(time.Second),
			Timeout:   xtime.Duration(time.Second),
			KeepAlive: xtime.Duration(time.Second * 10),
			Breaker: &breaker.Config{
				Window:  xtime.Duration(time.Second),
				Sleep:   xtime.Duration(time.Millisecond * 100),
				Bucket:  10,
				Ratio:   0.5,
				Request: 100,
			},
		},
		MaHTTPClient: &bm.ClientConfig{
			App: &bm.App{
				Key:    "f6433799dbd88751",
				Secret: "36f8ddb1806207fe07013ab6a77a3935",
			},
			Dial:      xtime.Duration(time.Second),
			Timeout:   xtime.Duration(time.Second),
			KeepAlive: xtime.Duration(time.Second * 10),
			Breaker: &breaker.Config{
				Window:  xtime.Duration(time.Second),
				Sleep:   xtime.Duration(time.Millisecond * 100),
				Bucket:  10,
				Ratio:   0.5,
				Request: 100,
			},
		},
		Session: &permit.SessionConfig{
			SessionIDLength: 32,
			CookieLifeTime:  1800,
			Domain:          ".bilibili.co",
		},
		ManagerHost:     "http://uat-manager.bilibili.co",
		DashboardHost:   "http://uat-auth-mng.bilibili.co",
		DashboardCaller: "manager-go",
	})


	e := bm.NewServer(nil)

	//Check whether the user has logged in
	e.GET("/login", a.Verify(), func(c *bm.Context) {
		c.JSON("pass", nil)
	})
	//Check whether the user has logged in,and check th user has the access permisson to the specifed path
	e.GET("/tag/del", a.Permit("TAG_DEL"), func(c *bm.Context) {
		uid := metadata.Int64(c, metadata.Uid)
		username := metadata.String(c, metadata.Username)
		c.JSON(fmt.Sprintf("pass uid(%d) username(%s)", uid, username), nil)
	})

	e.Run(":18080")
}
