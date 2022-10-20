package uctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
)

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
	}
	return ""
}

func SetSpanContext(ctx *gin.Context, spanCtx context.Context) {
	ctx.Set(consts.CtxSpanCtxVar, spanCtx)
}

func GetSpanContext(ctx context.Context) context.Context {
	c, ok := ctx.(*gin.Context)
	if ok {
		value, ok := c.Get(consts.CtxSpanCtxVar)
		if ok {
			c, ok := value.(context.Context)
			if ok {
				return c
			}
		}
	}
	return ctx
}
