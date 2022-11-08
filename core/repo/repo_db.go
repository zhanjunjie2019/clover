package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/uorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type RepoDBFactory struct {
	Repos         []defs.IRepo                            `allimpls:""`
	AppDefs       []defs.IAppDef                          `allimpls:""`
	DBFactory     uorm.DBFactoryIOCInterface              `singleton:""`
	OpenTelemetry opentelemetry.OpenTelemetryIOCInterface `singleton:""`
}

func (r *RepoDBFactory) Initialization() error {
	ctx, span := r.OpenTelemetry.Start(context.Background(), "Init DB Tables")
	defer span.End()

	db, err := r.DBFactory.GetDB()
	if err != nil {
		return err
	}
	ctx = uctx.SetAppDB(ctx, db)
	for _, repo := range r.Repos {
		err = repo.AutoMigrate(ctx)
		if err != nil {
			return err
		}
	}
	for _, app := range r.AppDefs {
		app.SetGormDB(db)
	}
	return nil
}
