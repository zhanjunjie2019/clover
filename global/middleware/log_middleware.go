package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"go.uber.org/zap"
	"io"
	"net/http"
	"runtime"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type LoggerMiddleware struct{}

func (l *LoggerMiddleware) MiddlewareWrapHandler() server.HandlerWrapper {
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			start := time.Now()
			layout := uctx.GetLogLayout(ctx)
			defer func() {
				// 异常日志处理
				if recoverErr := recover(); recoverErr != nil {
					buf := make([]byte, 1<<12)
					runtime.Stack(buf, false)
					layout.Error("请求故障：" + string(buf))
				}
				layout.AppendLogsFields(
					zap.Int64("rt", time.Since(start).Milliseconds()),
				)
				layout.Println()
			}()
			body := req.Body()
			bodyBts, err := json.Marshal(body)
			if err != nil {
				layout.Error("请求故障(请求体读取失败)：" + err.Error())
				return err
			}
			md, b := metadata.FromContext(ctx)
			if !b {
				layout.Error("请求故障(请求元数据读取失败)")
				return fmt.Errorf("请求故障(请求元数据读取失败)")
			}
			mdBts, err := json.Marshal(md)
			if err != nil {
				layout.Error("请求故障(请求元数据读取失败)：" + err.Error())
				return err
			}
			ctx, traceID := uctx.GetTraceIDByGrpcCtx(ctx)
			ctx, tenantID := uctx.GetTenantIDByGrpcCtx(ctx)
			layout.AppendLogsFields(
				zap.String("method", "gRPC"),
				zap.String("uri", req.Endpoint()),
				zap.String("reqBody", string(bodyBts)),
				zap.String("reqMD", string(mdBts)),
				zap.String("traceID", traceID),
				zap.String("tenantID", tenantID),
			)
			return handlerFunc(ctx, req, rsp)
		}
	}
}

func (l *LoggerMiddleware) MiddlewareHandlerFunc(option *defs.ControllerOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		layout := uctx.GetLogLayout(c)
		defer func() {
			// 异常日志处理
			if recoverErr := recover(); recoverErr != nil {
				buf := make([]byte, 1<<12)
				runtime.Stack(buf, false)
				layout.Error("请求故障：" + string(buf))
			}
			layout.AppendLogsFields(
				zap.Int("httpStatus", c.Writer.Status()),
				zap.Int64("rt", time.Since(start).Milliseconds()),
			)
			layout.Println()
		}()
		// 读取请求体
		var body, err = c.GetRawData()
		if err != nil {
			layout.Error("请求故障(请求体读取失败)：" + err.Error())
			return
		}
		// 请求头
		headerJson, err := l.getRequestHeaderJson(c.Request)
		if err != nil {
			layout.Error("请求故障(请求头读取失败)：" + err.Error())
			return
		}

		// 将原body塞回去
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		layout.AppendLogsFields(
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.RequestURI),
			zap.String("qry", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("reqBody", string(body)),
			zap.String("reqHeader", string(headerJson)),
			zap.String("traceID", uctx.GetTraceID(c)),
			zap.String("rUri", option.RelativePath),
			zap.String("tenantID", uctx.GetTenantID(c)),
		)
		c.Next()
	}
}

func (l LoggerMiddleware) getRequestHeaderJson(req *http.Request) ([]byte, error) {
	header := req.Header
	bs, err := json.Marshal(header)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
