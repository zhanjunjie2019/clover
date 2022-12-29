package defs

import "go-micro.dev/v4/server"

type IGrpcHandler interface {
	// GrpcRegister 注册服务
	GrpcRegister(s server.Server) error
}
