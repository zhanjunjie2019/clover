package uctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
	"go-micro.dev/v4/metadata"
)

func WithValueTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, consts.CtxTraceIDVar, traceID)
}

func SetTraceID(ctx *gin.Context, traceID string) {
	ctx.Set(consts.CtxTraceIDVar, traceID)
}

func GetTraceID(ctx context.Context) string {
	c, ok := ctx.(*gin.Context)
	if ok {
		value, ok := c.Get(consts.CtxTraceIDVar)
		if ok {
			s, ok := value.(string)
			if ok {
				return s
			}
		}
	} else {
		value := ctx.Value(consts.CtxTraceIDVar)
		if value != nil {
			s, ok := value.(string)
			if ok {
				return s
			}
		}
	}
	return ""
}

func GetTraceIDByGrpcCtx(ctx context.Context) (context.Context, string) {
	value := ctx.Value(consts.CtxTraceIDVar)
	if value != nil {
		s, ok := value.(string)
		if ok {
			return ctx, s
		}
	} else {
		traceID, ok := metadata.Get(ctx, consts.TraceIDHeaderKey)
		if ok && len(traceID) > 0 {
			ctx = WithValueTraceID(ctx, traceID)
			return ctx, traceID
		}
	}
	return ctx, ""
}

func GetTraceSpanIDByGrpcCtx(ctx context.Context) string {
	traceSpanID, ok := metadata.Get(ctx, consts.TraceSpanIDHeaderKey)
	if ok && len(traceSpanID) > 0 {
		return traceSpanID
	}
	return ""
}

func SetSpanContext(ctx *gin.Context, spanCtx context.Context) {
	ctx.Request = ctx.Request.WithContext(spanCtx)
}

func GetSpanContext(ctx context.Context) context.Context {
	c, ok := ctx.(*gin.Context)
	if ok {
		return c.Request.Context()
	}
	return ctx
}
