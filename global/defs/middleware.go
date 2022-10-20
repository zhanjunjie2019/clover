package defs

import "github.com/gin-gonic/gin"

type IMiddleware interface {
	MiddlewareHandlerFunc(option *ControllerOption) gin.HandlerFunc
}
