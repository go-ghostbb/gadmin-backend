package role

import (
	"context"
	"gadmin-backend/internal/model/entity"
	"gadmin-backend/internal/service"
	"gadmin-backend/utility/errorx"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"net/http"

	"gadmin-backend/api/role/v1"
)

func (c *ControllerV1) GetRoleByID(ctx context.Context, req *v1.GetRoleByIDReq) (res *v1.GetRoleByIDRes, err error) {
	var (
		role *entity.Role
		r    = g.RequestFromCtx(ctx)
	)
	role, err = service.Role().GetById(ctx, req.Id)
	if err != nil {
		if gerror.Is(err, errorx.ErrRecordNotFound) {
			return
		} else {
			// 除了record not found以外的錯誤都回傳500
			r.Response.Status = http.StatusInternalServerError
			return
		}
	}

	// convert
	if err = gconv.Struct(role, &res); err != nil {
		r.Response.Status = http.StatusInternalServerError
		return
	}

	return
}
