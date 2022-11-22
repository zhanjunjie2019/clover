package convs

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-example/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-example/bc/infr/repo/po"
)

func Entity1POToDO(entity1 po.ExampleEntity1) model.Entity1 {
	return model.NewEntity1(entity1.ID, model.Entity1Value{
		FirstName: entity1.FirstName,
		LastName:  entity1.LastName,
	})
}

func Entity1WithEntity2POToDO(entity1 po.ExampleEntity1, entity2 po.ExampleEntity2) model.Entity1 {
	do := Entity1POToDO(entity1)
	do.SetEntity2(Entity2POToDO(entity2))
	return do
}

func Entity1DOToPO(do model.Entity1) po.ExampleEntity1 {
	value := do.FullValue()
	return po.ExampleEntity1{
		ModelPO: defs.ModelPO{
			ID: do.ID(),
		},
		FirstName: value.FirstName,
		LastName:  value.LastName,
	}
}

func Entity1DOToPOWithEntity2(do model.Entity1) (po.ExampleEntity1, po.ExampleEntity2) {
	return Entity1DOToPO(do), Entity2DOToPO(do.GetEntity2())
}
