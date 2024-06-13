// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gadmin-backend/internal/model"
	"gadmin-backend/internal/model/entity"
)

type (
	IUser interface {
		// GetUserByUsernamePassword 根據使用者名稱和密碼獲取使用者資訊
		GetUserByUsernamePassword(ctx context.Context, in model.UserLoginInput) (user *entity.User, err error)
		// GetUserInfo 獲取使用者資訊
		GetUserInfo(ctx context.Context) (out *model.UserInfoOutput)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
