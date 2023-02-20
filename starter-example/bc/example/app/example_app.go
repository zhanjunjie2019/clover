package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/app/cmd"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/domain/gateway"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/domain/model"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/configs"
	_ "github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/gatewayimpl"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type ExampleApp struct {
	ExampleGateway gateway.IExampleGateway `singleton:"github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/gatewayimpl.ExampleGateway"`
}

func (e *ExampleApp) ExampleHellowWorld(ctx context.Context, c cmd.HelloWorldCmd) (rs cmd.HelloWorldResult, err error) {
	// 本服务的数据库事务方式
	err = uctx.AppTransaction(ctx, func(ctx context.Context) (err error) {
		entity := model.NewExampleEntity(0, model.ExampleEntityValue{
			FirstName: c.FirstName,
			LastName:  c.LastName,
		})
		entity.SetValueObject(model.ExampleValueObject{
			RandomValue1: utils.UUID(),
			RandomValue2: utils.UUID(),
			RandomValue3: utils.UUID(),
		})
		_, err = e.ExampleGateway.SaveExampleEntity(ctx, entity)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		// 发送异步通知
		err = e.ExampleGateway.PublishEventMessage(ctx, entity)
		if err != nil {
			err = errs.ToUnifiedError(err)
			return
		}
		// 读取动态配置
		exampleConfig := configs.GetExampleConfig()
		rs.Greetings = "Hello " + c.FirstName + " " + c.LastName + "!>" + exampleConfig.Aa.Bb
		return nil
	})
	return
}
