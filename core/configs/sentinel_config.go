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

type SentinelConfigDefine struct{}

func (s *SentinelConfigDefine) GetOption() defs.ConfigOptions {
	return defs.NewConfigOptions(
		defs.ConfigKey(coreconsts.SentinelConfigDefineKey),
		defs.ConfigFileName(coreconsts.SentinelConfigFileName),
		defs.ConfigCanLoadByConsul(true),
	)
}

func (s *SentinelConfigDefine) ReloadConfig(config any) error {
	sentinelConfig := config.(*confs.SentinelConfig)
	confs.SetSentinelConfig(*sentinelConfig)
	return nil
}

func (s *SentinelConfigDefine) Unmarshal(data []byte) (any, error) {
	var sc confs.SentinelConfig
	err := yaml.Unmarshal(data, &sc)
	return &sc, err
}
