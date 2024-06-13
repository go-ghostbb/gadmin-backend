package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// insert
// update
// delete

type GetRoleByIDReq struct {
	g.Meta `path:"/{id}" tags:"role" method:"get" summary:"根據id搜尋role" security:"BearerAuth"`
	Id     uint `json:"id"`
}

type GetRoleByIDRes struct {
	g.Meta      `mime:"application/json"`
	Id          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRoleReq struct {
	g.Meta      `path:"/" tags:"role" method:"post" summary:"創建角色" security:"BearerAuth"`
	Code        string `json:"code" v:"required"`
	Name        string `json:"name" v:"required"`
	Description string `json:"description"`
}

type CreateRoleRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteRoleReq struct {
	g.Meta `path:"/{id}" tags:"role" method:"delete" summary:"刪除角色" security:"BearerAuth"`
	Id     uint `json:"id"`
}

type DeleteRoleRes struct {
	g.Meta `mime:"application/json"`
}
