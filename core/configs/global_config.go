package configs

import (
	"github.com/zhanjunjie2019/clover/core/coreconsts"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/nsqd"
	"github.com/zhanjunjie2019/clover/global/redisc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IConfigDefine

type GlobalConfigDefine struct {
	RedisClient redisc.RedisClientIOCInterface `singleton:""`
	NsqProducer nsqd.NsqProducerIOCInterface   `singleton:""`
}

func (g *GlobalConfigDefine) GetOption() defs.ConfigOption {
	return defs.ConfigOption{
		ConfigKey:       coreconsts.GlobalConfigDefineKey,
		CanLoadByConsul: true,
		ConfigFileName:  coreconsts.GlobalConfigFileName,
	}
}

func (g *GlobalConfigDefine) ReloadConfig(config any) (err error) {
	globalConfig := config.(*confs.GlobalConfig)
	confs.SetGlobalConfig(*globalConfig)
	svcConf := confs.GetServerConfig().SvcConf
	var change bool
	// 热更redis配置
	redisConfig := globalConfig.RedisConfig
	if redisConfig.Enabled.Bool() {
		layout := defs.NewLogLayout(zapcore.InfoLevel, svcConf.SvcMode.Uint8(), svcConf.SvcName, svcConf.SvcNum, svcConf.SvcVersion)
		change, err = g.RedisClient.Create(redisConfig.Addr, redisConfig.Password, redisConfig.DB)
		if err != nil {
			layout.Error("Redis初始化失败!", zap.Error(err))
			layout.Println()
			return
		} else if change {
			layout.Info("Redis初始化成功!")
			layout.Println()
		}
	}
	// 热更nsq配置
	nsqConfig := globalConfig.NsqConfig
	if nsqConfig.Enabled.Bool() && len(nsqConfig.ProducerAddr) > 0 {
		layout := defs.NewLogLayout(zapcore.InfoLevel, svcConf.SvcMode.Uint8(), svcConf.SvcName, svcConf.SvcNum, svcConf.SvcVersion)
		change, err = g.NsqProducer.CreatePublisher(nsqConfig.ProducerAddr)
		if err != nil {
			layout.Error("NSQ.Producer初始化失败!", zap.Error(err))
			layout.Println()
			return
		} else if change {
			layout.Info("NSQ.Producer初始化成功!")
			layout.Println()
		}
	}
	return err
}

func (g *GlobalConfigDefine) Unmarshal(data []byte) (any, error) {
	var gc confs.GlobalConfig
	err := yaml.Unmarshal(data, &gc)
	return &gc, err
}
