package po

import "github.com/zhanjunjie2019/clover/global/defs"

type ExampleEntity1 struct {
	defs.ModelPO
	// FirstName 姓
	FirstName string `gorm:"column:first_name;comment:姓;size:64;"`
	// LastName 名
	LastName string `gorm:"column:last_name;comment:名;size:64;"`
}
