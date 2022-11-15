//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package consumer

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &server_{}
		},
	})
	serverStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Server{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(serverStructDescriptor)
}

type server_ struct {
	ConsumersStart_ func() error
}

func (s *server_) ConsumersStart() error {
	return s.ConsumersStart_()
}

type ServerIOCInterface interface {
	ConsumersStart() error
}

var _serverSDID string

func GetServerSingleton() (*Server, error) {
	if _serverSDID == "" {
		_serverSDID = util.GetSDIDByStructPtr(new(Server))
	}
	i, err := singleton.GetImpl(_serverSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Server)
	return impl, nil
}

func GetServerIOCInterfaceSingleton() (ServerIOCInterface, error) {
	if _serverSDID == "" {
		_serverSDID = util.GetSDIDByStructPtr(new(Server))
	}
	i, err := singleton.GetImplWithProxy(_serverSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServerIOCInterface)
	return impl, nil
}

type ThisServer struct {
}

func (t *ThisServer) This() ServerIOCInterface {
	thisPtr, _ := GetServerIOCInterfaceSingleton()
	return thisPtr
}
