package middleware

import (
	"context"
	sentapi "github.com/alibaba/sentinel-golang/api"
	sbase "github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"go-micro.dev/v4/server"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type SentinelMiddleware struct{}

func (s *SentinelMiddleware) MiddlewareWrapHandler() server.HandlerWrapper {
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx, tenantID := uctx.GetTenantIDByGrpcCtx(ctx)
			err := s.filter(tenantID, "")
			if err != nil {
				return err
			}
			return handlerFunc(ctx, req, rsp)
		}
	}
}

// MiddlewareWrapHandlerBySentinelStrategy 有需要的gRPC接口，要手动在注册时加上
func (s *SentinelMiddleware) MiddlewareWrapHandlerBySentinelStrategy(sentinelStrategy string) server.HandlerWrapper {
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx, tenantID := uctx.GetTenantIDByGrpcCtx(ctx)
			err := s.filter(tenantID, "-"+sentinelStrategy)
			if err != nil {
				return err
			}
			return handlerFunc(ctx, req, rsp)
		}
	}
}

func (s *SentinelMiddleware) MiddlewareHandlerFunc(option *defs.ControllerOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sentinelResourceSuffix string
		if option != nil {
			sentinelResourceSuffix = "-" + option.SentinelStrategy
		}
		err := s.filter(uctx.GetTenantID(c), sentinelResourceSuffix)
		if err != nil {
			response.FailWithMessage(c, err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func (s *SentinelMiddleware) filter(tenantID, sentinelResourceSuffix string) error {
	resource := confs.GetServerConfig().SvcConf.SvcName + sentinelResourceSuffix
	// 限流关键选项
	options := []sentapi.EntryOption{
		sentapi.WithTrafficType(sbase.Inbound),
		sentapi.WithResourceType(sbase.ResTypeWeb),
		// 热点参数，默认租户ID级别限流
		sentapi.WithAttachment(consts.TenantIDField, tenantID),
	}
	// 执行限流器
	entry, blockError := sentapi.Entry(
		resource,
		options...,
	)
	// 限流结果
	if blockError != nil {
		return errs.CurrentLimitingErr
	}
	defer entry.Exit()
	return nil
}
