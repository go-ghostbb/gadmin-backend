// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gadmin-backend/utility/claims"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IAuth interface {
		// MiddlewareFunc 中間件
		MiddlewareFunc(r *ghttp.Request)
		// GetPayload 獲取payload
		GetPayload(ctx context.Context) *claims.Access
		// GetIdentityKey 獲取Identity
		GetIdentityKey(ctx context.Context) uint
		// Login 登入
		Login(ctx context.Context) (accessTokenStr, refreshTokenStr string, expire time.Time)
		// Logout 登出
		Logout(ctx context.Context)
		// Refresh 刷新
		Refresh(ctx context.Context, oldAccessToken, oldRefreshToken string) (accessTokenStr, refreshTokenStr string, expire time.Time)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
