package convs

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-example/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-example/bc/infr/repo/po"
)

func Entity2POToDO(entity2 po.ExampleEntity2) model.Entity2 {
	return model.NewEntity2(entity2.ID, model.Entity2Value{
		RandomValue: entity2.RandomValue,
	})
}

func Entity2DOToPO(do model.Entity2) po.ExampleEntity2 {
	value := do.FullValue()
	return po.ExampleEntity2{
		ModelPO: defs.ModelPO{
			ID: do.ID(),
		},
		RandomValue: value.RandomValue,
	}
}
