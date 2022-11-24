package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/share/example/protobuf"
	"github.com/zhanjunjie2019/clover/starter-example/bc/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-example/bc/domain/model"
	"github.com/zhanjunjie2019/clover/starter-example/bc/infr/configs"
	_ "github.com/zhanjunjie2019/clover/starter-example/bc/infr/gatewayimpl"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IAppDef

type ExampleApp struct {
	ExampleGateway gateway.IExampleGateway `singleton:"github.com/zhanjunjie2019/clover/starter-example/bc/infr/gatewayimpl.ExampleGateway"`
	DB             *gorm.DB
}

func (e *ExampleApp) SetGormDB(db *gorm.DB) {
	e.DB = db
}

func (e *ExampleApp) ExampleHellowWord(ctx context.Context, firstName, lastName string) (greetings string, err error) {
	// 本服务的数据库事务方式,不需要的话：直接使用`ctx = uctx.WithValueAppDB(ctx, e.DB)`就可以了
	// 开启事务的tx，传到gateway，还可以开启下级子事务，树形嵌套
	err = e.DB.Transaction(func(tx *gorm.DB) (err error) {
		ctx = uctx.WithValueAppDB(ctx, tx)
		entity1 := model.NewExampleEntity(0, model.ExampleEntityValue{
			FirstName: firstName,
			LastName:  lastName,
		})
		entity1.SetValueObject(model.ExampleValueObject{
			RandomValue1: utils.UUID(),
			RandomValue2: utils.UUID(),
			RandomValue3: utils.UUID(),
		})
		_, err = e.ExampleGateway.SaveExampleEntity(ctx, entity1)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		// 读取动态配置
		exampleConfig := configs.GetExampleConfig()
		greetings = "Hello " + firstName + " " + lastName + "!>" + exampleConfig.Aa.Bb
		err = e.ExampleGateway.PublishEventMessage(ctx, protobuf.ExampleDTO{
			FirstName: firstName,
			LastName:  lastName,
		})
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		return nil
	})
	return
}
