package v1

import "github.com/gogf/gf/v2/frame/g"

type UserInfoReq struct {
	g.Meta `path:"/info" tags:"system" method:"get" summary:"使用者資訊" security:"BearerAuth"`
}

type UserInfoRes struct {
	g.Meta   `mime:"application/json"`
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}
