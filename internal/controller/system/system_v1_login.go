package system

import (
	"context"
	"gadmin-backend/api/system/v1"
	"gadmin-backend/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = new(v1.LoginRes)
	res.AccessToken, res.RefreshToken, res.Expires = service.Auth().Login(ctx)
	return
}
