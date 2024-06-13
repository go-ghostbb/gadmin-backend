package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"system" method:"post" summary:"登入"`
	Username string `json:"username" v:"required"`
	Password string `json:"password" v:"required"`
}
type LoginRes struct {
	g.Meta       `mime:"application/json"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Expires      time.Time `json:"expires"`
}
