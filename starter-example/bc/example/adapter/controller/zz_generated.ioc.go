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
			return &helloWorldController_{}
		},
	})
	helloWorldControllerStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &HelloWorldController{}
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
	allimpls.RegisterStructDescriptor(helloWorldControllerStructDescriptor)
}

type helloWorldController_ struct {
	GetOption_ func() defs.ControllerOption
	Handle_    func(c *gin.Context)
}

func (h *helloWorldController_) GetOption() defs.ControllerOption {
	return h.GetOption_()
}

func (h *helloWorldController_) Handle(c *gin.Context) {
	h.Handle_(c)
}

type HelloWorldControllerIOCInterface interface {
	GetOption() defs.ControllerOption
	Handle(c *gin.Context)
}

var _helloWorldControllerSDID string