package role

import (
	"context"
	"gadmin-backend/internal/dao"
	"gadmin-backend/internal/model"
	"gadmin-backend/internal/model/do"
	"gadmin-backend/internal/model/entity"
	"gadmin-backend/internal/service"
	"gadmin-backend/utility/errorx"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() service.IRole {
	return &sRole{}
}

// GetById 根據id獲取role
func (s *sRole) GetById(ctx context.Context, id uint) (role *entity.Role, err error) {
	var (
		qRole  = dao.Role.Ctx(ctx)
		record gdb.Record
	)
	record, err = qRole.Where(do.Role{Id: id}).One()
	if err != nil {
		return
	}
	if record.IsEmpty() {
		return nil, errorx.ErrRecordNotFound
	}

	if err = record.Struct(&role); err != nil {
		return
	}

	return
}

// Create 創建角色
func (s *sRole) Create(ctx context.Context, in model.CreateRoleInput) (err error) {
	var (
		qRole = dao.Role.Ctx(ctx)
	)

	_, err = qRole.Insert(do.Role{
		Code:        in.Code,
		Name:        in.Name,
		Description: in.Description,
	})
	return err
}

// Delete 刪除角色
func (s *sRole) Delete(ctx context.Context, id uint) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// role
		if _, err = tx.Model(do.Role{}).Where(do.Role{Id: id}).Delete(); err != nil {
			return err
		}

		// m2m
		if _, err = tx.Model(do.UserRole{}).Where(do.UserRole{RoleId: id}).Delete(); err != nil {
			return err
		}

		// commit
		return nil
	})
}
