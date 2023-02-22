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

type ServerConfigDefine struct{}

func (s *ServerConfigDefine) GetOption() defs.ConfigOptions {
	return defs.NewConfigOptions(
		defs.ConfigKey(coreconsts.ServerConfigDefineKey),
		defs.ConfigFileName(coreconsts.ServerConfigFileName),
	)
}

func (s *ServerConfigDefine) ReloadConfig(config any) error {
	serverConfig := config.(*confs.ServerConfig)
	confs.SetServerConfig(*serverConfig)
	return nil
}

func (s *ServerConfigDefine) Unmarshal(data []byte) (any, error) {
	var sc confs.ServerConfig
	err := yaml.Unmarshal(data, &sc)
	return &sc, err
}
