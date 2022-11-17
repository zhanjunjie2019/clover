package app

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/domain/gateway"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IAppDef

type UserApp struct {
	UserGateway gateway.IUserGateway `singleton:"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/gatewayimpl.UserGateway"`
	DB          *gorm.DB
}

func (u *UserApp) SetGormDB(db *gorm.DB) {
	u.DB = db
}

func (u *UserApp) UserAuthorizationCode(ctx context.Context, userName, password, redirect string) (authorizationCode, redirectUrl string, err error) {
	ctx = uctx.WithValueAppDB(ctx, u.DB)
	// TODO
	return
}
