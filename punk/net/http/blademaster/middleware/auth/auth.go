package auth

import (
	"angrymiao-go/punk/ecode"
	jt "angrymiao-go/punk/jwt"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/metadata"
	"github.com/dgrijalva/jwt-go"
)
const (
	_authorization = "Authorization"
)

// Auth is the authorization middleware
type Auth struct {
}

// authFunc will return mid and error by given context
type authFunc func(*bm.Context) (int64, string, error)

// New is used to create an authorization middleware
func New() *Auth {
	auth := &Auth{}
	return auth
}

// User is used to mark path as access required.
// If `access_key` is exist in request form, it will using mobile access policy.
// Otherwise to web access policy.
func (a *Auth) User(ctx *bm.Context) {
	a.UserMobile(ctx)
}

// UserMobile is used to mark path as mobile access required.
func (a *Auth) UserMobile(ctx *bm.Context) {
	a.midAuth(ctx, a.AuthToken)
}

// Guest is used to mark path as guest policy.
// If `access_key` is exist in request form, it will using mobile access policy.
// Otherwise to web access policy.
func (a *Auth) Guest(ctx *bm.Context) {
	a.GuestMobile(ctx)
}

// GuestMobile is used to mark path as mobile guest policy.
func (a *Auth) GuestMobile(ctx *bm.Context) {
	a.guestAuth(ctx, a.AuthToken)
}

// AuthToken is used to authorize request by token
func (a *Auth) AuthToken(ctx *bm.Context) (int64, string,  error) {
	req := ctx.Request
	accessToken := req.Header.Get(_authorization)
	if accessToken == "" {
		return 0, "",  ecode.NoLogin
	}

	var code ecode.Code
	claims, err := jt.ParseToken(accessToken)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			code = ecode.NoLogin
		default:
			code = ecode.NoLogin
		}
	}
	if code != ecode.OK {
		ctx.JSON(nil, code)
		ctx.Abort()
		return 0, "", ecode.NoLogin
	}
	return claims.UserID, claims.Phone, nil
}

func (a *Auth) midAuth(ctx *bm.Context, auth authFunc) {
	mid, username, err := auth(ctx)
	if err != nil {
		ctx.JSON(nil, err)
		ctx.Abort()
		return
	}
	setMid(ctx, mid, username)
}

func (a *Auth) guestAuth(ctx *bm.Context, auth authFunc) {
	mid, username, err := auth(ctx)
	// no error happened and mid is valid
	if err == nil && mid > 0 {
		setMid(ctx, mid, username)
		return
	}
	return
}

// set mid into context
// NOTE: This method is not thread safe.
func setMid(ctx *bm.Context, mid int64, username string) {
	ctx.Set("mid", mid)
	if md, ok := metadata.FromContext(ctx); ok {
		md[metadata.Mid] = mid
		md[metadata.Username] = username
		return
	}
}
