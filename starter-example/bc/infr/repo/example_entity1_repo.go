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

type ExampleEntity1Repo struct {
}

func (e *ExampleEntity1Repo) AutoMigrate(ctx context.Context) error {
	return uctx.GetAppDBWithCtx(ctx).AutoMigrate()
}

func (e *ExampleEntity1Repo) Save(ctx context.Context, entity1 po.ExampleEntity1) (id defs.ID, err error) {
	err = uctx.GetAppDBWithCtx(ctx).Save(&entity1).Error
	return entity1.ID, err
}
