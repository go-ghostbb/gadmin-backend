package model

type UserLoginInput struct {
	Username string
	Password string
}

type UserInfoOutput struct {
	Id       uint
	Username string
}
