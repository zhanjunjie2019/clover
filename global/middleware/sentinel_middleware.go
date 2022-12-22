package middleware

import (
	sentapi "github.com/alibaba/sentinel-golang/api"
	sbase "github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type SentinelMiddleware struct{}

func (s *SentinelMiddleware) MiddlewareHandlerFunc(option *defs.ControllerOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 必须限流已开启，并且全局限流，或接口拥有指定的限流策略
		if confs.GetSentinelConfig().Enabled == 1 {
			if option == nil || len(option.SentinelStrategy) > 0 {
				resource := confs.GetServerConfig().SvcConf.SvcName
				if option != nil {
					resource = resource + option.GetSentinelResourceSuffix()
				}
				// 热点参数
				attrs := map[any]any{
					consts.ClientIpField: c.ClientIP(),
					consts.TenantIDField: uctx.GetTenantID(c),
				}
				// 限流关键选项
				options := []sentapi.EntryOption{
					sentapi.WithTrafficType(sbase.Inbound),
					sentapi.WithResourceType(sbase.ResTypeWeb),
					sentapi.WithAttachments(attrs),
				}
				// 执行限流器
				entry, blockError := sentapi.Entry(
					resource,
					options...,
				)
				// 限流结果
				if blockError != nil {
					response.FailWithMessage(c, errs.CurrentLimitingErr)
					c.Abort()
					return
				}
				defer entry.Exit()
			}
		}
		c.Next()
	}
}
