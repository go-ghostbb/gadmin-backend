// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"gadmin-backend/api/system/v1"
)

type ISystemV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error)
}
