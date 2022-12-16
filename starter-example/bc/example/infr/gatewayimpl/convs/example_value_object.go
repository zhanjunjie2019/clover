package convs

import (
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/domain/model"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/repo/po"
)

func ValueObjectPOToDO(valueObject po.ExampleValueObject) model.ExampleValueObject {
	return model.ExampleValueObject{
		RandomValue1: valueObject.RandomValue1,
		RandomValue2: valueObject.RandomValue2,
		RandomValue3: valueObject.RandomValue3,
	}
}

func ValueObjectDOToPO(do model.ExampleValueObject) po.ExampleValueObject {
	return po.ExampleValueObject{
		RandomValue1: do.RandomValue1,
		RandomValue2: do.RandomValue2,
		RandomValue3: do.RandomValue3,
	}
}
