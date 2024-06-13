package v1

import "github.com/gogf/gf/v2/frame/g"

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"system" method:"delete" summary:"登出"`
}

type LogoutRes struct {
	g.Meta `mime:"application/json"`
}
