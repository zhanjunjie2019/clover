package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"go-micro.dev/v4/server"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap/zapcore"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type TraceMiddleware struct {
	OpenTelemetry opentelemetry.OpenTelemetryIOCInterface `singleton:""`
}

func (t *TraceMiddleware) MiddlewareWrapHandler() server.HandlerWrapper {
	otelConfig := confs.GetGlobalConfig().OtelConfig
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			if otelConfig.Enabled == 1 {
				ctx = t.OpenTelemetry.ExtractByGrpcCtx(ctx)
				var span trace.Span
				ctx, span = t.OpenTelemetry.Start(ctx, "GRPC "+req.Endpoint())
				defer span.End()
				// 链路上下文中传递日志输出器
				layout := defs.NewLogLayout(zapcore.InfoLevel)
				ctx = uctx.WithValueLogLayout(ctx, layout)
				// 获取链路追踪ID
				var traceID = span.SpanContext().TraceID().String()
				ctx = uctx.WithValueTraceID(ctx, traceID)
			} else {
				// 构造链路ID
				var traceID = utils.UUID()
				ctx = uctx.WithValueTraceID(ctx, traceID)
			}
			return handlerFunc(ctx, req, rsp)
		}
	}
}

func (t *TraceMiddleware) MiddlewareHandlerFunc(option *defs.ControllerOptions) gin.HandlerFunc {
	otelConfig := confs.GetGlobalConfig().OtelConfig
	return func(c *gin.Context) {
		if otelConfig.Enabled == 1 {
			req := c.Request
			ctx := t.OpenTelemetry.Extract(req.Context(), req.Header)
			ctx, span := t.OpenTelemetry.Start(ctx, "HTTP "+req.Method+" "+req.RequestURI)
			defer span.End()
			// 需要传递链路上下文
			tenantID := uctx.GetTenantID(c)
			if len(tenantID) > 0 {
				ctx = uctx.WithValueTenantID(ctx, tenantID)
			}
			// 链路上下文中传递日志输出器
			layout := defs.NewLogLayout(zapcore.InfoLevel)
			uctx.SetLogLayout2GinCtx(c, layout)
			ctx = uctx.WithValueLogLayout(ctx, layout)
			// 传递链路上下文
			uctx.SetSpanContext(c, ctx)
			// 获取链路追踪ID
			var traceID = span.SpanContext().TraceID().String()
			uctx.SetTraceID(c, traceID)
		} else {
			// 构造链路ID
			var traceID = utils.UUID()
			uctx.SetTraceID(c, traceID)
		}
		c.Next()
	}
}
