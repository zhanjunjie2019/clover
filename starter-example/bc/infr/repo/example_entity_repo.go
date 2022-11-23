package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-example/bc/infr/repo/po"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IRepo

type ExampleEntityRepo struct {
}

func (e *ExampleEntityRepo) AutoMigrate(ctx context.Context) error {
	return uctx.GetAppDBWithCtx(ctx).AutoMigrate(&po.ExampleEntity{})
}

func (e *ExampleEntityRepo) Save(ctx context.Context, entity po.ExampleEntity) (id defs.ID, err error) {
	err = uctx.GetAppDBWithCtx(ctx).Save(&entity).Error
	return entity.ID, err
}
