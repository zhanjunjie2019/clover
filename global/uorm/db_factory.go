package uorm

import (
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"gorm.io/gorm"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type DBFactory struct {
	globalDB *gorm.DB
}

func (d *DBFactory) GetDB() (*gorm.DB, error) {
	if d.globalDB == nil {
		db, err := createDb(confs.GetDBConfig())
		if err != nil {
			return nil, err
		}
		err = db.Use(otelgorm.NewPlugin())
		if err != nil {
			return nil, err
		}
		d.globalDB = db
	}
	return d.globalDB, nil
}

func createDb(linkConf confs.DBConfig) (*gorm.DB, error) {
	if linkConf.Postgres != nil {
		return gormPgSql(*linkConf.Postgres)
	} else if linkConf.Mysql != nil {
		return gormMysql(*linkConf.Mysql)
	}
	return nil, errs.DBLinkTypeErr
}
