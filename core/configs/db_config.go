package configs

import (
	"github.com/zhanjunjie2019/clover/core/coreconsts"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uorm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IConfigDefine

type DBConfigDefine struct {
	DBFactory uorm.DBFactoryIOCInterface `singleton:""`
}

func (d *DBConfigDefine) GetOption() defs.ConfigOption {
	return defs.ConfigOption{
		ConfigKey:       coreconsts.DBConfigDefineKey,
		CanLoadByConsul: true,
		ConfigFileName:  coreconsts.DBConfigFileName,
	}
}

func (d *DBConfigDefine) ReloadConfig(config any) error {
	dbConfig := config.(*confs.DBConfig)
	// 无法实现从无到有或从有到无的数据库连接热更
	if dbConfig.Enabled.Bool() {
		svcConf := confs.GetServerConfig().SvcConf
		layout := defs.NewLogLayout(zapcore.InfoLevel, svcConf.SvcMode.Uint8(), svcConf.SvcName, svcConf.SvcNum, svcConf.SvcVersion)
		change, err := d.DBFactory.Create(*dbConfig)
		if err != nil {
			layout.Error("数据库初始化失败!", zap.Error(err))
			layout.Println()
			return err
		} else if change {
			layout.Info("数据库初始化成功!")
			layout.Println()
		}
	}
	return nil
}

func (d *DBConfigDefine) Unmarshal(data []byte) (any, error) {
	var dc confs.DBConfig
	err := yaml.Unmarshal(data, &dc)
	return &dc, err
}
