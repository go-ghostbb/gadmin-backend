// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"gadmin-backend/api/role/v1"
)

type IRoleV1 interface {
	GetRoleByID(ctx context.Context, req *v1.GetRoleByIDReq) (res *v1.GetRoleByIDRes, err error)
	CreateRole(ctx context.Context, req *v1.CreateRoleReq) (res *v1.CreateRoleRes, err error)
	DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error)
}
