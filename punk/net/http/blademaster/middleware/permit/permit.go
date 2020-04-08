package permit

import (
	"angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/net/rpc/warden"
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"

	mng "angrymiao-go/app/admin/manager/api"
	"angrymiao-go/punk/ecode"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/metadata"
)

const (
	_permissionURI         = "/x/admin/manager/permission"
	_sessIDKey             = "_AJSESSIONID"
	_sessUIDKey            = "uid"      // manager user_id
	_sessUnKey             = "username" // LDAP username
	_defaultDomain         = ".angrymiao.com"
	_defaultCookieLifeTime = 2592000
	// CtxPermissions will be set into ctx.
	CtxPermissions = "permissions"
)

// permissions .
type permissions struct {
	UID   int64    `json:"uid"`
	Perms []string `json:"perms"`
}

// Permit is an auth middleware.
type Permit struct {
	permissionURI   string
	dashboardCaller string
	maClient        *bm.Client // manager-admin client

	sm *SessionManager // user Session

	mng.PermitClient // mng grpc client
}

//Verify only export Verify function because of less configure
type Verify interface {
	Verify() bm.HandlerFunc
}

// Config identify config.
type Config struct {
	DsHTTPClient    *bm.ClientConfig // auth client config. appkey can not reuse.
	MaHTTPClient    *bm.ClientConfig // manager-admin client config
	Session         *SessionConfig
	ManagerHost     string
}


// New new an auth service.
func New(c *Config) *Permit {
	a := &Permit{
		permissionURI:   c.ManagerHost + _permissionURI,
		maClient:        bm.NewClient(c.MaHTTPClient),
		sm:              newSessionManager(c.Session),
	}
	return a
}

func New2(c *warden.ClientConfig) *Permit {
	permitClient, err := mng.NewClient(c)
	if err != nil {
		panic(errors.WithMessage(err, "Failed to dial mng rpc server"))
	}
	return &Permit{
		PermitClient: permitClient,
		sm:           &SessionManager{},
	}
}


// Verify return bm HandlerFunc which check whether the user has logged in.
func (p *Permit) Verify() bm.HandlerFunc {
	return func(ctx *bm.Context) {
		si, err := p.login(ctx)
		if err != nil {
			ctx.JSON(nil, ecode.Unauthorized)
			ctx.Abort()
			return
		}
		p.sm.SessionRelease(ctx, si)
	}
}

// Permit return bm HandlerFunc which check whether the user has logged in and has the access permission of the location.
// If `permit` is empty,it will allow any logged in request.
func (p *Permit) Permit(permit string) bm.HandlerFunc {
	return func(ctx *bm.Context) {
		var (
			si    *Session
			uid   int64
			perms []string
			err   error
		)
		si, err = p.login(ctx)
		if err != nil {
			ctx.JSON(nil, ecode.Unauthorized)
			ctx.Abort()
			return
		}
		defer p.sm.SessionRelease(ctx, si)
		uid, perms, err = p.permissions(ctx, si.Get(_sessUnKey).(string))
		if err == nil {
			si.Set(_sessUIDKey, uid)
			ctx.Set(_sessUIDKey, uid)
			if md, ok := metadata.FromContext(ctx); ok {
				md[metadata.Uid] = uid
			}
		}
		if len(perms) > 0 {
			ctx.Set(CtxPermissions, perms)
		}
		if !p.permit(permit, perms) {
			ctx.JSON(nil, ecode.AccessDenied)
			ctx.Abort()
			return
		}
	}
}

// login check whether the user has logged in.
func (p *Permit) login(ctx *bm.Context) (si *Session, err error) {
	si = p.sm.SessionStart(ctx)
	if si.Get(_sessUnKey) == nil {
		var username string
		if username, err = p.verify(ctx); err != nil {
			return
		}
		si.Set(_sessUnKey, username)
	}
	ctx.Set(_sessUnKey, si.Get(_sessUnKey))
	if md, ok := metadata.FromContext(ctx); ok {
		md[metadata.Username] = si.Get(_sessUnKey)
	}
	return
}


func (p *Permit) verify(ctx *bm.Context) (username string, err error) {
	var (
		sid string
		r   = ctx.Request
	)
	session, err := r.Cookie(_sessIDKey)
	if err == nil {
		sid = session.Value
	}
	if sid == "" {
		err = ecode.Unauthorized
		return
	}
	username, err = p.verifyByCache(ctx, sid)
	return
}

// permit check whether user has the access permission of the location.
func (p *Permit) permit(permit string, permissions []string) bool {
	if permit == "" {
		return true
	}
	for _, p := range permissions {
		if p == permit {
			// access the permit
			return true
		}
	}
	return false
}

// verifyByCache check whether the user is valid from Dashboard.
func (p *Permit) verifyByCache(ctx *bm.Context, sid string) (username string, err error) {
	var dsbRes struct {
		UserName string `json:"username"`
	}
	conn := p.sm.rp.Get(ctx)
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("GET", sid))
	if err != nil {
		log.Error("conn.Do(GET, %s) error(%v)", sid, err)
		return
	}
	if err = json.Unmarshal(reply, &dsbRes); err != nil {
		log.Error("json.Unmarshal(%v,%v) error(%v)", string(reply), dsbRes, err)
	}

	username = dsbRes.UserName
	return
}

// permissions get user's permisssions from manager-admin.
//从permission service获取相应的权限
func (p *Permit) permissions(ctx *bm.Context, username string) (uid int64, perms []string, err error) {
	params := url.Values{}
	//TODO 根据username获取permission，暂时不处理
	params.Set(_sessUnKey, username)
	var res struct {
		Code int         `json:"code"`
		Data permissions `json:"data"`
	}
	if err = p.maClient.Get(ctx, p.permissionURI, metadata.String(ctx, metadata.RemoteIP), params, &res); err != nil {
		log.Error("auth get permissions url(%s) error(%v)", p.permissionURI+"?"+params.Encode(), err)
		return
	}
	if ecode.Int(res.Code) != ecode.OK {
		log.Error("auth get permissions url(%s) error(%v)", p.permissionURI+"?"+params.Encode(), res.Code)
		err = ecode.Int(res.Code)
		return
	}
	perms = res.Data.Perms
	uid = res.Data.UID
	return
}
