// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} // 主鍵
	CreatedAt *gtime.Time // 創建時間
	UpdatedAt *gtime.Time // 更新時間
	Username  interface{} // 使用者名稱
	Password  interface{} // 密碼
}
