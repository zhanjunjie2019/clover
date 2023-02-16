//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package config

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
	"github.com/zhanjunjie2019/clover/global/defs"
	timex "time"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configDefines_{}
		},
	})
	configDefinesStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ConfigDefines{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IScheduler),
					},
				},
			},
		},
	}
	singleton.RegisterStructDescriptor(configDefinesStructDescriptor)
	allimpls.RegisterStructDescriptor(configDefinesStructDescriptor)
}

type configDefines_ struct {
	GetTaskTypeCode_       func() string
	GetSpec_               func() string
	GetLockDuration_       func() timex.Duration
	RunTask_               func(ctx contextx.Context) error
	LoadAllConfigByLocal_  func() error
	LoadAllConfigByConsul_ func() error
	loadConfigByLocal_     func(configDefine defs.IConfigDefine) (any, error)
	loadConfigByConsul_    func(configDefine defs.IConfigDefine) (any, error)
	newConfigLoader_       func() error
}

func (c *configDefines_) GetTaskTypeCode() string {
	return c.GetTaskTypeCode_()
}

func (c *configDefines_) GetSpec() string {
	return c.GetSpec_()
}

func (c *configDefines_) GetLockDuration() timex.Duration {
	return c.GetLockDuration_()
}

func (c *configDefines_) RunTask(ctx contextx.Context) error {
	return c.RunTask_(ctx)
}

func (c *configDefines_) LoadAllConfigByLocal() error {
	return c.LoadAllConfigByLocal_()
}

func (c *configDefines_) LoadAllConfigByConsul() error {
	return c.LoadAllConfigByConsul_()
}

func (c *configDefines_) loadConfigByLocal(configDefine defs.IConfigDefine) (any, error) {
	return c.loadConfigByLocal_(configDefine)
}

func (c *configDefines_) loadConfigByConsul(configDefine defs.IConfigDefine) (any, error) {
	return c.loadConfigByConsul_(configDefine)
}

func (c *configDefines_) newConfigLoader() error {
	return c.newConfigLoader_()
}

type ConfigDefinesIOCInterface interface {
	GetTaskTypeCode() string
	GetSpec() string
	GetLockDuration() timex.Duration
	RunTask(ctx contextx.Context) error
	LoadAllConfigByLocal() error
	LoadAllConfigByConsul() error
	loadConfigByLocal(configDefine defs.IConfigDefine) (any, error)
	loadConfigByConsul(configDefine defs.IConfigDefine) (any, error)
	newConfigLoader() error
}

var _configDefinesSDID string

func GetConfigDefinesSingleton() (*ConfigDefines, error) {
	if _configDefinesSDID == "" {
		_configDefinesSDID = util.GetSDIDByStructPtr(new(ConfigDefines))
	}
	i, err := singleton.GetImpl(_configDefinesSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ConfigDefines)
	return impl, nil
}

func GetConfigDefinesIOCInterfaceSingleton() (ConfigDefinesIOCInterface, error) {
	if _configDefinesSDID == "" {
		_configDefinesSDID = util.GetSDIDByStructPtr(new(ConfigDefines))
	}
	i, err := singleton.GetImplWithProxy(_configDefinesSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ConfigDefinesIOCInterface)
	return impl, nil
}

type ThisConfigDefines struct {
}

func (t *ThisConfigDefines) This() ConfigDefinesIOCInterface {
	thisPtr, _ := GetConfigDefinesIOCInterfaceSingleton()
	return thisPtr
}
