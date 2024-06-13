package role

import (
	"context"
	"gadmin-backend/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"

	"gadmin-backend/api/role/v1"
)

func (c *ControllerV1) DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error) {
	var (
		r = g.RequestFromCtx(ctx)
	)
	err = service.Role().Delete(ctx, req.Id)
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
	}
	return nil, err
}
