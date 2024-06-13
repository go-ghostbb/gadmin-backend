package role

import (
	"context"
	"gadmin-backend/internal/model"
	"gadmin-backend/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"

	"gadmin-backend/api/role/v1"
)

func (c *ControllerV1) CreateRole(ctx context.Context, req *v1.CreateRoleReq) (res *v1.CreateRoleRes, err error) {
	var (
		r = g.RequestFromCtx(ctx)
	)
	err = service.Role().Create(ctx, model.CreateRoleInput{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		r.Response.Status = http.StatusInternalServerError
	}
	return nil, err
}
