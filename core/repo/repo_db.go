package repo

import (
	"context"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	"github.com/zhanjunjie2019/clover/global/uorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type RepoDBFactory struct {
	Repos         []defs.IRepo                            `allimpls:""`
	DBFactory     uorm.DBFactoryIOCInterface              `singleton:""`
	OpenTelemetry opentelemetry.OpenTelemetryIOCInterface `singleton:""`
}

func (r *RepoDBFactory) Initialization() error {
	db, err := r.DBFactory.GetDB()
	if err != nil {
		return err
	}
	ctx, span := r.OpenTelemetry.Start(context.Background(), "Init DB Tables")
	defer span.End()
	for _, repo := range r.Repos {
		err := repo.AutoMigrate(ctx, db)
		if err != nil {
			return err
		}
	}
	return nil
}
