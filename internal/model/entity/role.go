// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id          uint        `json:"id"          orm:"id"          ` // 主鍵
	CreatedAt   *gtime.Time `json:"created_at"  orm:"created_at"  ` // 創建時間
	UpdatedAt   *gtime.Time `json:"updated_at"  orm:"updated_at"  ` // 更新時間
	Code        string      `json:"code"        orm:"code"        ` // 角色代碼
	Name        string      `json:"name"        orm:"name"        ` // 角色名稱
	Description string      `json:"description" orm:"description" ` // 說明
}
