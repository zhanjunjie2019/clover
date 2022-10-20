package uorm

import (
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func gormPgSql(dbc confs.GeneralDBConf) (*gorm.DB, error) {
	pgsqlConfig := postgres.Config{
		DSN:                  pgSqlDsn(dbc),
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), config()); err != nil {
		return nil, errs.DBConnectionErr
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbc.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbc.MaxOpenConns)
		return db, nil
	}
}

func pgSqlDsn(p confs.GeneralDBConf) string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
}
