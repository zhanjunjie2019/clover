package defs

import "go-micro.dev/v4/server"

type IGrpcServiceHandler interface {
	// GrpcRegister 注册服务
	GrpcRegister(s server.Server) error
}
