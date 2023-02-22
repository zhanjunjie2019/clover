//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package configs

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
	"github.com/zhanjunjie2019/clover/global/defs"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &dBConfigDefine_{}
		},
	})
	dBConfigDefineStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &DBConfigDefine{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IConfigDefine),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(dBConfigDefineStructDescriptor)
	var _ defs.IConfigDefine = &DBConfigDefine{}
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &globalConfigDefine_{}
		},
	})
	globalConfigDefineStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &GlobalConfigDefine{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IConfigDefine),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(globalConfigDefineStructDescriptor)
	var _ defs.IConfigDefine = &GlobalConfigDefine{}
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &sentinelConfigDefine_{}
		},
	})
	sentinelConfigDefineStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &SentinelConfigDefine{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IConfigDefine),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(sentinelConfigDefineStructDescriptor)
	var _ defs.IConfigDefine = &SentinelConfigDefine{}
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serverConfigDefine_{}
		},
	})
	serverConfigDefineStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ServerConfigDefine{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IConfigDefine),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(serverConfigDefineStructDescriptor)
	var _ defs.IConfigDefine = &ServerConfigDefine{}
}

type dBConfigDefine_ struct {
	GetOption_    func() defs.ConfigOptions
	ReloadConfig_ func(config any) error
	Unmarshal_    func(data []byte) (any, error)
}

func (d *dBConfigDefine_) GetOption() defs.ConfigOptions {
	return d.GetOption_()
}

func (d *dBConfigDefine_) ReloadConfig(config any) error {
	return d.ReloadConfig_(config)
}

func (d *dBConfigDefine_) Unmarshal(data []byte) (any, error) {
	return d.Unmarshal_(data)
}

type globalConfigDefine_ struct {
	GetOption_    func() defs.ConfigOptions
	ReloadConfig_ func(config any) (err error)
	Unmarshal_    func(data []byte) (any, error)
}

func (g *globalConfigDefine_) GetOption() defs.ConfigOptions {
	return g.GetOption_()
}

func (g *globalConfigDefine_) ReloadConfig(config any) (err error) {
	return g.ReloadConfig_(config)
}

func (g *globalConfigDefine_) Unmarshal(data []byte) (any, error) {
	return g.Unmarshal_(data)
}

type sentinelConfigDefine_ struct {
	GetOption_    func() defs.ConfigOptions
	ReloadConfig_ func(config any) error
	Unmarshal_    func(data []byte) (any, error)
}

func (s *sentinelConfigDefine_) GetOption() defs.ConfigOptions {
	return s.GetOption_()
}

func (s *sentinelConfigDefine_) ReloadConfig(config any) error {
	return s.ReloadConfig_(config)
}

func (s *sentinelConfigDefine_) Unmarshal(data []byte) (any, error) {
	return s.Unmarshal_(data)
}

type serverConfigDefine_ struct {
	GetOption_    func() defs.ConfigOptions
	ReloadConfig_ func(config any) error
	Unmarshal_    func(data []byte) (any, error)
}

func (s *serverConfigDefine_) GetOption() defs.ConfigOptions {
	return s.GetOption_()
}

func (s *serverConfigDefine_) ReloadConfig(config any) error {
	return s.ReloadConfig_(config)
}

func (s *serverConfigDefine_) Unmarshal(data []byte) (any, error) {
	return s.Unmarshal_(data)
}

type DBConfigDefineIOCInterface interface {
	GetOption() defs.ConfigOptions
	ReloadConfig(config any) error
	Unmarshal(data []byte) (any, error)
}

type GlobalConfigDefineIOCInterface interface {
	GetOption() defs.ConfigOptions
	ReloadConfig(config any) (err error)
	Unmarshal(data []byte) (any, error)
}

type SentinelConfigDefineIOCInterface interface {
	GetOption() defs.ConfigOptions
	ReloadConfig(config any) error
	Unmarshal(data []byte) (any, error)
}

type ServerConfigDefineIOCInterface interface {
	GetOption() defs.ConfigOptions
	ReloadConfig(config any) error
	Unmarshal(data []byte) (any, error)
}

var _dBConfigDefineSDID string
var _globalConfigDefineSDID string
var _sentinelConfigDefineSDID string
var _serverConfigDefineSDID string
