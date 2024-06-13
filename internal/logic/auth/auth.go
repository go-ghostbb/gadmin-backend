package auth

import (
	"context"
	"gadmin-backend/internal/model"
	"gadmin-backend/internal/model/entity"
	"gadmin-backend/internal/service"
	"gadmin-backend/utility/claims"
	"gadmin-backend/utility/errorx"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type sAuth struct {
	// signing algorithm - possible values are HS256, HS384, HS512, RS256, RS384 or RS512
	// Optional, default is HS256.
	SigningAlgorithm string

	// Secret key used for signing. Required.
	Key []byte

	// TokenHeadName is a string in the header. Default value is "Bearer"
	TokenHeadName string

	// Duration that a jwt token is valid. Optional, defaults to one hour.
	AccessTimeout time.Duration

	// RefreshTimeout
	RefreshTimeout time.Duration

	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	AccessTimeFunc func() time.Time

	// RefreshTimeFunc
	RefreshTimeFunc func() time.Time

	// BlacklistPrefix
	BlacklistPrefix string
}

var (
	// TokenKey default jwt token key in params
	TokenKey = "JWT_TOKEN"
	// PayloadKey default jwt payload key in params
	PayloadKey = "JWT_PAYLOAD"
	// IdentityKey default identity key
	IdentityKey = "identity"
	// The blacklist stores tokens that have not expired but have been deactivated.
	blacklist = gcache.New()
)

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	svc := &sAuth{
		SigningAlgorithm: "HS256",
		Key:              []byte("secret key"),
		AccessTimeout:    time.Minute * 10,
		AccessTimeFunc:   time.Now,
		RefreshTimeout:   time.Hour * 72,
		RefreshTimeFunc:  time.Now,
		TokenHeadName:    "Bearer",
		BlacklistPrefix:  "jwt:blacklist:",
	}
	return svc
}

// 驗證
func (s *sAuth) authenticator(ctx context.Context) (user *entity.User, err error) {
	var (
		req = g.RequestFromCtx(ctx)
		in  model.UserLoginInput
	)

	if err = req.Parse(&in); err != nil {
		return
	}

	user, err = service.User().GetUserByUsernamePassword(ctx, in)
	if err != nil {
		return
	}

	return
}

// payload
func (s *sAuth) payloadFunc(user *entity.User) *claims.Access {
	return &claims.Access{
		Id:       user.Id,
		Username: user.Username,
	}
}

// 驗證失敗
func (s *sAuth) unauthorized(ctx context.Context, status int, err error) {
	var (
		req = g.RequestFromCtx(ctx)
	)
	req.Response.Status = status

	req.Response.WriteJson(g.Map{
		"code":    status,
		"message": err.Error(),
	})

	req.ExitAll()
}

// signed
func (s *sAuth) signedString(token *jwt.Token) (tokenStr string, err error) {
	return token.SignedString(s.Key)
}

// 從header獲取jwt
func (s *sAuth) jwtFromHeader(r *ghttp.Request, key string) (string, error) {
	authHeader := r.Header.Get(key)

	if authHeader == "" {
		return "", errorx.ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)

	if (len(parts) != 2) || parts[0] != s.TokenHeadName {
		return "", errorx.ErrInvalidAuthHeader
	}

	return parts[1], nil
}

// 解析token
func (s *sAuth) parseAccessToken(r *ghttp.Request) (*jwt.Token, error) {
	token, err := s.jwtFromHeader(r, "Authorization")
	if err != nil {
		return nil, err
	}

	return jwt.ParseWithClaims(token, &claims.Access{}, func(token *jwt.Token) (interface{}, error) {
		return s.Key, nil
	})
}

// 是否在黑名單
func (s *sAuth) inBlacklist(ctx context.Context, token string) (bool, error) {
	tokenRaw, err := gmd5.EncryptString(token)
	if err != nil {
		return false, nil
	}

	key := s.BlacklistPrefix + tokenRaw

	if in, err := blacklist.Contains(ctx, key); err != nil {
		return false, nil
	} else {
		return in, nil
	}
}

// 放入黑名單
func (s *sAuth) setBlacklist(ctx context.Context, token string, claimsAcc *claims.Access) error {
	tokenRaw, err := gmd5.EncryptString(token)
	if err != nil {
		return nil
	}

	key := s.BlacklistPrefix + tokenRaw

	// (過期時間 + 緩存時間) - 現在時間
	// 緩存時間暫定為一小時
	duration := claimsAcc.ExpiresAt.Add(time.Hour).Sub(s.AccessTimeFunc()).Truncate(time.Second)

	err = blacklist.Set(ctx, key, true, duration)
	return err
}

// 從jwt獲取claims
func (s *sAuth) getClaimsFromJWT(ctx context.Context) (*claims.Access, string, error) {
	r := g.RequestFromCtx(ctx)

	token, err := s.parseAccessToken(r)
	if err != nil {
		return nil, "", errorx.ErrInvalidToken
	}

	return token.Claims.(*claims.Access), token.Raw, nil
}

// MiddlewareFunc 中間件
func (s *sAuth) MiddlewareFunc(r *ghttp.Request) {
	var (
		ctx       = r.GetCtx()
		claimsAcc *claims.Access
		token     string
		err       error
	)

	claimsAcc, token, err = s.getClaimsFromJWT(ctx)
	if err != nil {
		s.unauthorized(ctx, http.StatusUnauthorized, err)
		return
	}

	in, err := s.inBlacklist(ctx, token)
	if err != nil {
		s.unauthorized(ctx, http.StatusUnauthorized, err)
		return
	}

	if in {
		s.unauthorized(ctx, http.StatusUnauthorized, errorx.ErrInvalidToken)
		return
	}

	r.SetParam(TokenKey, token)
	r.SetParam(PayloadKey, claimsAcc)
	r.SetParam(IdentityKey, claimsAcc.Id)
}

// GetPayload 獲取payload
func (s *sAuth) GetPayload(ctx context.Context) *claims.Access {
	var (
		r = g.RequestFromCtx(ctx)
	)
	return r.GetParam(PayloadKey).Interface().(*claims.Access)
}

// GetIdentityKey 獲取Identity
func (s *sAuth) GetIdentityKey(ctx context.Context) uint {
	var (
		r = g.RequestFromCtx(ctx)
	)
	return r.GetParam(IdentityKey).Uint()
}

// Login 登入
func (s *sAuth) Login(ctx context.Context) (accessTokenStr, refreshTokenStr string, expire time.Time) {
	user, err := s.authenticator(ctx)
	if err != nil {
		if gerror.Is(err, errorx.ErrFailedAuthentication) {
			s.unauthorized(ctx, http.StatusUnauthorized, err)
			return
		} else {
			s.unauthorized(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	// accessToken 跟 refreshToken 對照key
	key := guid.S()

	// create the accessToken
	expire = s.AccessTimeFunc().Add(s.AccessTimeout)

	var (
		accessToken  = jwt.New(jwt.GetSigningMethod(s.SigningAlgorithm))
		accessClaims = s.payloadFunc(user)
	)
	accessClaims.Issuer = "gadmin"
	accessClaims.ExpiresAt = jwt.NewNumericDate(expire)
	accessClaims.IssuedAt = jwt.NewNumericDate(s.AccessTimeFunc())
	accessClaims.Key = key

	accessToken.Claims = accessClaims

	accessTokenStr, err = s.signedString(accessToken)
	if err != nil {
		s.unauthorized(ctx, http.StatusUnauthorized, errorx.ErrFailedTokenCreation)
		return
	}

	// create the refreshToken
	refreshExpire := s.RefreshTimeFunc().Add(s.RefreshTimeout)

	var (
		refreshToken  = jwt.New(jwt.GetSigningMethod(s.SigningAlgorithm))
		refreshClaims = claims.Refresh{Key: key}
	)
	refreshClaims.Issuer = "gadmin"
	accessClaims.ExpiresAt = jwt.NewNumericDate(refreshExpire)
	accessClaims.IssuedAt = jwt.NewNumericDate(s.RefreshTimeFunc())

	refreshToken.Claims = refreshClaims

	refreshTokenStr, err = s.signedString(refreshToken)
	if err != nil {
		s.unauthorized(ctx, http.StatusUnauthorized, errorx.ErrFailedTokenCreation)
		return
	}

	return
}

// Logout 登出
func (s *sAuth) Logout(ctx context.Context) {
	var (
		r         = g.RequestFromCtx(ctx)
		token     = r.GetParam(TokenKey).String()
		claimsAcc = r.GetParam(PayloadKey).Interface().(*claims.Access)
		err       error
	)

	err = s.setBlacklist(ctx, token, claimsAcc)
	if err != nil {
		s.unauthorized(ctx, http.StatusInternalServerError, err)
	}
}
