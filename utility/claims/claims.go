package claims

import "github.com/golang-jwt/jwt/v5"

type Access struct {
	jwt.RegisteredClaims
	Key      string `json:"key"`
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type Refresh struct {
	jwt.RegisteredClaims
	Key string `json:"key"`
}
