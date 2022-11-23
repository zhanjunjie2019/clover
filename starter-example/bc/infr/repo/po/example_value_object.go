package po

import "github.com/zhanjunjie2019/clover/global/defs"

type ExampleValueObject struct {
	defs.ModelPO
	// 实体ID
	EntityID defs.ID `gorm:"column:entity_id;comment:实体ID;"`
	// 随机数据1
	RandomValue1 string `gorm:"column:random_value1;comment:随机数据1;size:64;"`
	// 随机数据2
	RandomValue2 string `gorm:"column:random_value2;comment:随机数据2;size:64;"`
	// 随机数据3
	RandomValue3 string `gorm:"column:random_value3;comment:随机数据3;size:64;"`
}

func (e ExampleValueObject) Equats(t ExampleValueObject) bool {
	if e.EntityID != t.EntityID {
		return false
	}
	if e.RandomValue1 != t.RandomValue1 {
		return false
	}
	if e.RandomValue2 != t.RandomValue2 {
		return false
	}
	if e.RandomValue3 != t.RandomValue3 {
		return false
	}
	return true
}
