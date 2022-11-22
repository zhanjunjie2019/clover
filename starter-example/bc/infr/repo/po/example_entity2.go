package po

import "github.com/zhanjunjie2019/clover/global/defs"

type ExampleEntity2 struct {
	defs.ModelPO
	// 主表ID
	Entity1ID defs.ID `gorm:"column:entity_1_id;comment:主表ID;"`
	// 随机数据
	RandomValue string `gorm:"column:random_value;comment:随机数据;size:64;"`
}
