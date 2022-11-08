package defs

import "gorm.io/gorm"

type IAppDef interface {
	SetGormDB(db *gorm.DB)
}
