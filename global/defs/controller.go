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
	RelativePath string
	// HttpMethod 请求方法 使用 http.MethodGet 等相关枚举
	HttpMethod string
	// AuthCode 接口资源编码
	AuthCode string
	// Middlewares 自定义中间件
	Middlewares []IMiddleware
	// SentinelStrategy 熔断限流策略配置
	SentinelStrategy string
}

func (c ControllerOption) GetSentinelResourceSuffix() string {
	rs := ""
	if len(c.SentinelStrategy) > 0 {
		rs += "-" + c.SentinelStrategy
	}
	if len(c.AuthCode) > 0 {
		rs += "-" + c.AuthCode
	}
	return rs
}
