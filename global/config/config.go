package config

import (
	consulCfg "github.com/asim/go-micro/plugins/config/source/consul/v4"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/utils"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"os"
	"sync"
)

var enc = yaml.NewEncoder()

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type ConfigDefines struct {
	ConfigDefineCaches []defs.IConfigDefine `allimpls:""`
	rw                 sync.RWMutex
	configLoader       config.Config
}

func (d *ConfigDefines) LoadAllConfigByLocal() error {
	d.rw.Lock()
	defer d.rw.Unlock()
	for _, cache := range d.ConfigDefineCaches {
		if cache.GetOption().CanLoadByConsul {
			continue
		}
		conf, err := d.loadConfigByLocal(cache)
		if err != nil {
			return err
		}
		err = utils.LoadEnvToStruct(conf)
		if err != nil {
			return err
		}
		err = cache.ReloadConfig(conf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *ConfigDefines) LoadAllConfigByConsul() error {
	d.rw.Lock()
	defer d.rw.Unlock()
	for _, cache := range d.ConfigDefineCaches {
		if !cache.GetOption().CanLoadByConsul {
			continue
		}
		conf, err := d.loadConfigByLocal(cache)
		if err != nil {
			return err
		}
		configByConsul, err := d.loadConfigByConsul(cache)
		if err != nil {
			return err
		}
		err = utils.MergeStruct(conf, configByConsul)
		if err != nil {
			return err
		}
		err = utils.LoadEnvToStruct(conf)
		if err != nil {
			return err
		}
		err = cache.ReloadConfig(conf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *ConfigDefines) loadConfigByLocal(configDefine defs.IConfigDefine) (any, error) {
	content, err := os.ReadFile(consts.LocalConfigPathPre + configDefine.GetOption().ConfigFileName)
	if err != nil {
		return nil, err
	}
	conf, err := configDefine.Unmarshal(content)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (d *ConfigDefines) loadConfigByConsul(configDefine defs.IConfigDefine) (any, error) {
	if !configDefine.GetOption().CanLoadByConsul {
		return nil, nil
	}
	if d.configLoader == nil {
		err := d.newConfigLoader()
		if err != nil {
			return nil, err
		}
	}
	bytes := d.configLoader.Get(configDefine.GetOption().ConfigFileName).Bytes()
	obj, err := configDefine.Unmarshal(bytes)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func withYamlOption() source.Option {
	return func(o *source.Options) {
		o.Encoder = enc
	}
}

func (d *ConfigDefines) newConfigLoader() error {
	conf, err := config.NewConfig(
		config.WithReader(
			json.NewReader(
				reader.WithEncoder(enc),
			),
		),
	)
	if err != nil {
		return err
	}
	serverConfig := confs.GetServerConfig()
	consulAddr := serverConfig.ConsulConf.ConsulAddr
	svcName := serverConfig.SvcConf.SvcName
	version := serverConfig.SvcConf.SvcVersion
	// ?????????????????????
	sources := []source.Source{
		consulCfg.NewSource(
			consulCfg.WithAddress(consulAddr),
			consulCfg.WithPrefix("/"+svcName+"/default"),
			consulCfg.StripPrefix(true),
			withYamlOption(),
		),
		consulCfg.NewSource(
			consulCfg.WithAddress(consulAddr),
			consulCfg.WithPrefix("/"+svcName+"/"+version),
			consulCfg.StripPrefix(true),
			withYamlOption(),
		),
	}
	err = conf.Load(sources...)
	if err != nil {
		return err
	}
	d.configLoader = conf
	return nil
}
