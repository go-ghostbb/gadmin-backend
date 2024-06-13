package system

import (
	"context"
	"gadmin-backend/internal/service"

	"gadmin-backend/api/system/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	service.Auth().Logout(ctx)
	return
}
