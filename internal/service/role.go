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
	IRole interface {
		// GetById 根據id獲取role
		GetById(ctx context.Context, id uint) (role *entity.Role, err error)
		// Create 創建角色
		Create(ctx context.Context, in model.CreateRoleInput) (err error)
		Delete(ctx context.Context, id uint) (err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
