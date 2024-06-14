package system

import (
	"context"
	"gadmin-backend/internal/service"

	"gadmin-backend/api/system/v1"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	out := service.User().GetUserInfoByCtx(ctx)
	res = new(v1.UserInfoRes)
	res.Id = out.Id
	res.Username = out.Username
	res.Roles = out.Roles
	return
}
