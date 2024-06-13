// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta      `orm:"table:role, do:true"`
	Id          interface{} // 主鍵
	CreatedAt   *gtime.Time // 創建時間
	UpdatedAt   *gtime.Time // 更新時間
	Code        interface{} // 角色代碼
	Name        interface{} // 角色名稱
	Description interface{} // 說明
}
