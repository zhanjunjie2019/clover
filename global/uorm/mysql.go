package uorm

import (
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func gormMysql(dbc confs.GeneralDBConf) (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		DSN:                       mySqlDsn(dbc),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), config()); err != nil {
		return nil, errs.DBConnectionErr
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbc.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbc.MaxOpenConns)
		return db, nil
	}
}

func mySqlDsn(m confs.GeneralDBConf) string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
