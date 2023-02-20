package uorm

import (
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func gormMysql(dbc confs.DBConfig) (db *gorm.DB, err error) {
	mysqlConfig := mysql.Config{
		DSN:                       mySqlDsn(dbc),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), config()); err != nil {
		return nil, errs.DBConnectionErr
	}
	return db, updateMysqlConns(dbc, db)
}

func updateMysqlConns(dbc confs.DBConfig, db *gorm.DB) error {
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(dbc.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbc.MaxOpenConns)
		return nil
	} else {
		return errs.DBConnectionErr
	}
}

func mySqlDsn(m confs.DBConfig) string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
