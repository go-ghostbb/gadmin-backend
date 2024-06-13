// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"         orm:"id"         ` // 主鍵
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` // 創建時間
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` // 更新時間
	Username  string      `json:"username"   orm:"username"   ` // 使用者名稱
	Password  string      `json:"password"   orm:"password"   ` // 密碼
}
