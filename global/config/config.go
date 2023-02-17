package config

import (
	"context"
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
	"os"
	"sync"
	"time"
)

var enc = yaml.NewEncoder()

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IScheduler

type ConfigDefines struct {
	ConfigDefineCaches []defs.IConfigDefine `allimpls:""`
	rw                 sync.RWMutex
	configLoader       config.Config
}

// GetTaskTypeCode 定时任务
func (d *ConfigDefines) GetTaskTypeCode() string {
	return "config.ConfigDefineScheduler"
}

// GetSpec 每秒执行一次
func (d *ConfigDefines) GetSpec() string {
	return "* * * * * ?"
}

// GetLockDuration 不验证分布式锁
func (d *ConfigDefines) GetLockDuration() time.Duration {
	return 0
}

// RunTask 获取最新的配置
func (d *ConfigDefines) RunTask(ctx context.Context) error {
	return d.LoadAllConfigByConsul()
}

// LoadAllConfigByLocal 加载本地配置
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
	err = conf.Load(sources...)
	if err != nil {
		return err
	}
	d.configLoader = conf
	return nil
}
