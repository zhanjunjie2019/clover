package defs

import (
	"gorm.io/gorm"
	"strconv"
)

// ID 主键
type ID uint64

func (i ID) Uint() uint {
	return uint(i)
}

func (i ID) Int() int {
	return int(i)
}

func (i ID) Int64() int64 {
	return int64(i)
}

func (i ID) String() string {
	return strconv.Itoa(int(i))
}

type ModelPO struct {
	// 主键，默认自增长
	ID ID `json:"id" gorm:"primarykey"`
	// 乐观锁
	Revision uint64 `json:"-" gorm:"column:revision;comment:乐观锁;size:64;"`
	// 创建时间戳
	CreatedAt uint64 `json:"createdAt" gorm:"autoCreateTime:milli"`
	// 修改时间戳
	UpdatedAt uint64 `json:"updatedAt" gorm:"autoUpdateTime:milli"`
	// 删除标识符
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (m *ModelPO) BeforeCreate(tx *gorm.DB) (err error) {
	m.Revision += 1
	return
}

func (m *ModelPO) BeforeUpdate(tx *gorm.DB) (err error) {
	m.Revision += 1
	return
}
