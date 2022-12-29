package defs

import "github.com/gin-gonic/gin"

type IController interface {
	// GetOption 关键配置
	GetOption() ControllerOption
	// Handle 核心方法
	Handle(c *gin.Context)
}

// ControllerOption 接口配置
type ControllerOption struct {
	// RelativePath 接口路径
	RelativePath string `validate:"required"`
	// HttpMethod 请求方法 使用 http.MethodGet 等相关枚举
	HttpMethod string `validate:"required"`
	// AuthCodes 接口资源编码，拥有任意一个资源就算通过
	AuthCodes []string `validate:"dive,required"`
	// Middlewares 自定义中间件
	Middlewares []IHttpMiddleware
	// SentinelStrategy 熔断限流策略，不同得接口不要使用相同得策略
	SentinelStrategy string
}
