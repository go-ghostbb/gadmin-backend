package system

import (
	"context"
	"gadmin-backend/internal/service"

	"gadmin-backend/api/system/v1"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	res = new(v1.RefreshTokenRes)
	res.AccessToken, res.RefreshToken, res.Expires = service.Auth().Refresh(ctx, req.AccessToken, req.RefreshToken)
	return
}
