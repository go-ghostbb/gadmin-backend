package user

import (
	"context"
	"gadmin-backend/internal/dao"
	"gadmin-backend/internal/model"
	"gadmin-backend/internal/model/do"
	"gadmin-backend/internal/model/entity"
	"gadmin-backend/internal/service"
	"gadmin-backend/utility/errorx"
	"gadmin-backend/utility/password"
	"github.com/gogf/gf/v2/database/gdb"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// GetUserByUsernamePassword 根據使用者名稱和密碼獲取使用者資訊
func (s *sUser) GetUserByUsernamePassword(ctx context.Context, in model.UserLoginInput) (user *entity.User, err error) {
	var (
		qUser  = dao.User.Ctx(ctx)
		result gdb.Record
	)

	// 查詢使用者
	result, err = qUser.Where(do.User{Username: in.Username}).One()
	if err != nil {
		return
	}

	// 解析
	if err = result.Struct(&user); err != nil {
		return
	}

	// 確認密碼
	if result.IsEmpty() || !password.Check(user.Password, in.Password) {
		return nil, errorx.ErrFailedAuthentication
	}

	return
}

// GetUserInfoByCtx 從上下文獲取使用者資訊
func (s *sUser) GetUserInfoByCtx(ctx context.Context) (out *model.UserInfoOutput) {
	payload := service.Auth().GetPayload(ctx)
	out = new(model.UserInfoOutput)
	out.Id = payload.Id
	out.Username = payload.Username
	out.Roles = payload.Roles
	return out
}

// GetRoles 獲取角色列表
func (s *sUser) GetRoles(ctx context.Context, id uint) (roles []*entity.Role, err error) {
	var (
		qRole  = dao.Role
		qM2M   = dao.UserRole
		result gdb.Result
	)

	result, err = qRole.Ctx(ctx).
		LeftJoinOnFields(qM2M.Table(), qRole.Columns().Id, "=", qM2M.Columns().RoleId).
		Where(do.UserRole{UserId: id}).All()
	if err != nil {
		return nil, err
	}

	// convert
	if err = result.Structs(&roles); err != nil {
		return nil, err
	}

	return roles, nil
}
