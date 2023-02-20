package uorm

import (
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"gorm.io/gorm"
	"sync"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type DBFactory struct {
	dbConf   confs.DBConfig
	rw       sync.RWMutex
	globalDB *gorm.DB
}

func (d *DBFactory) Create(linkConf confs.DBConfig) (bool, error) {
	d.rw.Lock()
	defer d.rw.Unlock()
	// 完全一致，不用重新构建链接
	if d.dbConf.AllAaccordance(linkConf) {
		return false, nil
	}
	// 仅主要配置一致，也不用重新构建链接，但是需要重新设置连接数
	if d.dbConf.MainAaccordance(linkConf) {
		err := d.updateDb(linkConf)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	// 主要配置不一致时，需要重新构建链接
	db, err := d.createDb(linkConf)
	if err != nil {
		return false, err
	}
	err = db.Use(otelgorm.NewPlugin())
	if err != nil {
		return false, err
	}
	d.globalDB = db
	d.dbConf = linkConf
	return true, nil
}

func (d *DBFactory) GetDB() *gorm.DB {
	d.rw.RLock()
	defer d.rw.RUnlock()
	return d.globalDB
}

func (d *DBFactory) createDb(linkConf confs.DBConfig) (*gorm.DB, error) {
	if linkConf.DbType == confs.PostgreSQL {
		return gormPgSql(linkConf)
	} else if linkConf.DbType == confs.MySQL {
		return gormMysql(linkConf)
	}
	return nil, errs.DBLinkTypeErr
}

func (d *DBFactory) updateDb(linkConf confs.DBConfig) error {
	if linkConf.DbType == confs.PostgreSQL {
		return updatePgSqlConns(linkConf, d.globalDB)
	} else if linkConf.DbType == confs.MySQL {
		return updateMysqlConns(linkConf, d.globalDB)
	}
	return errs.DBLinkTypeErr
}
