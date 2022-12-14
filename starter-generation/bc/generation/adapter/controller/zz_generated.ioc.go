//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package controller

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &moduleGenController_{}
		},
	})
	moduleGenControllerStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ModuleGenController{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IController),
					},
				},
			},
		},
	}
	allimpls.RegisterStructDescriptor(moduleGenControllerStructDescriptor)
}

type moduleGenController_ struct {
	GetOption_ func() defs.ControllerOptions
	Handle_    func(c *gin.Context)
}

func (m *moduleGenController_) GetOption() defs.ControllerOptions {
	return m.GetOption_()
}

func (m *moduleGenController_) Handle(c *gin.Context) {
	m.Handle_(c)
}

type ModuleGenControllerIOCInterface interface {
	GetOption() defs.ControllerOptions
	Handle(c *gin.Context)
}

var _moduleGenControllerSDID string
