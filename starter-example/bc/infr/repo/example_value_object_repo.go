package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-example/bc/infr/repo/po"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IRepo

type ExampleValueObjectRepo struct {
}

func (e *ExampleValueObjectRepo) AutoMigrate(ctx context.Context) error {
	return uctx.GetAppDBWithCtx(ctx).AutoMigrate(&po.ExampleValueObject{})
}

func (e *ExampleValueObjectRepo) Save(ctx context.Context, valueObject po.ExampleValueObject) (id defs.ID, err error) {
	dbWithCtx := uctx.GetAppDBWithCtx(ctx)
	var val2 po.ExampleValueObject
	err = dbWithCtx.Where("entity_id=?", valueObject.EntityID).First(&val2).Error
	if err != gorm.ErrRecordNotFound {
		// 库中存在一致的值对象，则不改变
		if valueObject.Equats(val2) {
			return val2.ID, nil
		}
		// 库中存在不一致的值对象，则删除了再创建
		err = dbWithCtx.Delete(&val2).Error
		if err != nil {
			return
		}
	}
	err = dbWithCtx.Save(&valueObject).Error
	return valueObject.ID, err
}
