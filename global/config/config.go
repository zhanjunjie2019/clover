package config

import (
	"fmt"
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
	"go-micro.dev/v4/util/log"
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

// LoadAllConfigByLocal 加载本地配置
func (d *ConfigDefines) LoadAllConfigByLocal() error {
	d.rw.Lock()
	defer d.rw.Unlock()
	for _, cache := range d.ConfigDefineCaches {
		// 防止本地未加载的情况下，依赖的配置先加载了
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

// LoadAllConfigByConsul 加载配置中心配置
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
		err = d.loadConfigByConsul(cache, conf)
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

func (d *ConfigDefines) loadConfigByConsul(configDefine defs.IConfigDefine, conf any) error {
	if !configDefine.GetOption().CanLoadByConsul {
		return nil
	}
	if d.configLoader == nil {
		err := d.newConfigLoader()
		if err != nil {
			return err
		}
	}
	watcher, err := d.configLoader.Watch(configDefine.GetOption().ConfigFileName)
	if err != nil {
		return err
	}
	go func(conf any) {
		// 循环等待配置变化
		for {
			if v, err := watcher.Next(); err == nil {
				err = configMerge(configDefine, conf, v.Bytes())
				if err != nil {
					log.Error()
				}
			}
		}
	}(conf)

	bytes := d.configLoader.Get(configDefine.GetOption().ConfigFileName).Bytes()
	err = configMerge(configDefine, conf, bytes)
	return err
}

func configMerge(configDefine defs.IConfigDefine, conf any, consulByte []byte) (err error) {
	configByConsul, err := configDefine.Unmarshal(consulByte)
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
	err = configDefine.ReloadConfig(conf)
	if err != nil {
		return err
	}
	return err
}

func withYamlOption() source.Option {
	return func(o *source.Options) {
		o.Encoder = enc
	}
}

func (d *ConfigDefines) newConfigLoader() error {
	cfg, err := config.NewConfig(
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
	if len(consulAddr) == 0 {
		return fmt.Errorf("configuration node does not exist")
	}

	// 优先级从低到高
	var sources []source.Source
	for _, node := range serverConfig.ConsulConf.ConfigNode {
		sources = append(sources, consulCfg.NewSource(
			consulCfg.WithAddress(consulAddr),
			consulCfg.WithPrefix(node),
			consulCfg.StripPrefix(true),
			withYamlOption(),
		))
	}
	err = cfg.Load(sources...)
	if err != nil {
		return err
	}
	d.configLoader = cfg
	return nil
}
