package po

import "github.com/zhanjunjie2019/clover/global/defs"

type User struct {
	defs.ModelPO
	// TenantID 租户ID
	TenantID string `json:"tenantID" gorm:"column:tenant_id;comment:租户ID;size:64;"`
	// UserName 用户名
	UserName string `json:"userName" gorm:"column:user_name;comment:用户名;size:64;"`
	// Password 密码
	Password string `json:"password" gorm:"column:password;comment:密码;size:64;"`
}

func (u User) TableName() string {
	return "users_" + u.TenantID
}
