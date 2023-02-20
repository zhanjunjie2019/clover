package uorm

import (
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func gormPgSql(dbc confs.DBConfig) (db *gorm.DB, err error) {
	pgsqlConfig := postgres.Config{
		DSN:                  pgSqlDsn(dbc),
		PreferSimpleProtocol: false,
	}
	if db, err = gorm.Open(postgres.New(pgsqlConfig), config()); err != nil {
		return nil, errs.DBConnectionErr
	}
	return db, updatePgSqlConns(dbc, db)
}

func updatePgSqlConns(dbc confs.DBConfig, db *gorm.DB) error {
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(dbc.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbc.MaxOpenConns)
		return nil
	} else {
		return errs.DBConnectionErr
	}
}

func pgSqlDsn(p confs.DBConfig) string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
}
