package defs

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/server"
)

type IHttpMiddleware interface {
	MiddlewareHandlerFunc(option *ControllerOption) gin.HandlerFunc
}

type IGrpcMiddleware interface {
	MiddlewareWrapHandler() server.HandlerWrapper
}
