package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/opentelemetry"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type TraceMiddleware struct {
	OpenTelemetry opentelemetry.OpenTelemetryIOCInterface `singleton:""`
}

func (t *TraceMiddleware) MiddlewareHandlerFunc(option *defs.ControllerOption) gin.HandlerFunc {
	otelConfig := confs.GetGlobalConfig().OtelConfig
	return func(c *gin.Context) {
		if otelConfig.Enabled == 1 {
			req := c.Request
			ctx := t.OpenTelemetry.Extract(req.Context(), req.Header)
			ctx, span := t.OpenTelemetry.Start(ctx, "HTTP "+req.Method+" "+req.RequestURI)
			defer span.End()
			// 需要传递上下文
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
