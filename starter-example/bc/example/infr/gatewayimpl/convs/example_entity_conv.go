package convs

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/domain/model"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/repo/po"
)

func EntityPOToDO(entity po.ExampleEntity) model.ExampleEntity {
	return model.NewExampleEntity(entity.ID, model.ExampleEntityValue{
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	})
}

func EntityWithValueObjectPOToDO(entity po.ExampleEntity, valueObject po.ExampleValueObject) model.ExampleEntity {
	do := EntityPOToDO(entity)
	do.SetValueObject(ValueObjectPOToDO(valueObject))
	return do
}

func EntityDOToPO(do model.ExampleEntity) po.ExampleEntity {
	value := do.FullValue()
	return po.ExampleEntity{
		ModelPO: defs.ModelPO{
			ID: do.ID(),
		},
		FirstName: value.FirstName,
		LastName:  value.LastName,
	}
}

func EntityDOToPOWithValueObject(do model.ExampleEntity) (po.ExampleEntity, po.ExampleValueObject) {
	return EntityDOToPO(do), ValueObjectDOToPO(do.GetValueObject())
}
