package model

import "time"

type UserLoginInput struct {
	Username string
	Password string
}

type UserLoginOutput struct {
	Roles        []string
	AccessToken  string
	RefreshToken string
	Expire       time.Time
}

type UserInfoOutput struct {
	Id       uint
	Username string
	Roles    []string
}
