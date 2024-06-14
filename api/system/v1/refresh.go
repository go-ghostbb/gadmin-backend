package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type RefreshTokenReq struct {
	g.Meta       `path:"/refresh-token" tags:"system" method:"post" summary:"刷新token"`
	AccessToken  string `json:"accessToken" v:"required"`
	RefreshToken string `json:"refreshToken" v:"required"`
}

type RefreshTokenRes struct {
	g.Meta       `mime:"application/json"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Expires      time.Time `json:"expires"`
}
