package uorm

import "gorm.io/gorm"

func config() *gorm.Config {
	return &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
