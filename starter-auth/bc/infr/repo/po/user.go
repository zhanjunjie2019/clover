package po

import "github.com/zhanjunjie2019/clover/global/defs"

type User struct {
	defs.ModelPO
	// UserName 用户名
	UserName string `gorm:"column:user_name;comment:用户名;size:64;"`
	// Password 密码
	Password string `gorm:"column:password;comment:密码;size:64;"`
}
