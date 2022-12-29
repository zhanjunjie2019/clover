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
			return &authConfigDefine_{}
		},
	})
	authConfigDefineStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &AuthConfigDefine{}
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
	allimpls.RegisterStructDescriptor(authConfigDefineStructDescriptor)
}

type authConfigDefine_ struct {
	GetOption_    func() defs.ConfigOption
	ReloadConfig_ func(config any) error
	Unmarshal_    func(data []byte) (any, error)
}

func (a *authConfigDefine_) GetOption() defs.ConfigOption {
	return a.GetOption_()
}

func (a *authConfigDefine_) ReloadConfig(config any) error {
	return a.ReloadConfig_(config)
}

func (a *authConfigDefine_) Unmarshal(data []byte) (any, error) {
	return a.Unmarshal_(data)
}

type AuthConfigDefineIOCInterface interface {
	GetOption() defs.ConfigOption
	ReloadConfig(config any) error
	Unmarshal(data []byte) (any, error)
}

var _authConfigDefineSDID string