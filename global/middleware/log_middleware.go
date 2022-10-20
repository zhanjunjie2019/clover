package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
	"runtime"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type LoggerMiddleware struct{}

func (l *LoggerMiddleware) MiddlewareHandlerFunc(option *defs.ControllerOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		layout := defs.NewLogLayout(zapcore.InfoLevel)
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
		uctx.SetLogLayout(c, layout)
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
			zap.String("authCode", option.AuthCode),
			zap.String("rUri", option.RelativePath),
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
