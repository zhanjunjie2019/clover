package configs

import (
	"github.com/zhanjunjie2019/clover/core/coreconsts"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"gopkg.in/yaml.v3"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IConfigDefine

type DBConfigDefine struct{}

func (d *DBConfigDefine) GetOption() defs.ConfigOption {
	return defs.ConfigOption{
		ConfigKey:       coreconsts.DBConfigDefineKey,
		CanLoadByConsul: true,
		ConfigFileName:  coreconsts.DBConfigFileName,
	}
}

func (d *DBConfigDefine) ReloadConfig(config any) error {
	dbConfig := config.(*confs.DBConfig)
	confs.SetDBConfig(*dbConfig)
	return nil
}

func (d *DBConfigDefine) Unmarshal(data []byte) (any, error) {
	var dc confs.DBConfig
	err := yaml.Unmarshal(data, &dc)
	return &dc, err
}
