package uctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"go.uber.org/zap"
)

func SetLogLayout2GinCtx(ctx *gin.Context, layout *defs.LogLayout) {
	ctx.Set(consts.CtxLogsVar, layout)
}

func WithValueLogLayout(ctx context.Context, layout *defs.LogLayout) context.Context {
	return context.WithValue(ctx, consts.CtxLogsVar, layout)
}

func GetLogLayout(ctx context.Context) *defs.LogLayout {
	c, ok := ctx.(*gin.Context)
	if ok {
		value, ok := c.Get(consts.CtxLogsVar)
		if ok {
			layout, ok := value.(*defs.LogLayout)
			if ok {
				return layout
			}
		}
	} else {
		value := ctx.Value(consts.CtxLogsVar)
		if value != nil {
			layout, ok := value.(*defs.LogLayout)
			if ok {
				return layout
			}
		}
	}
	return nil
}

func AppendLogsFields(ctx context.Context, fields ...zap.Field) {
	layout := GetLogLayout(ctx)
	if layout != nil {
		layout.AppendLogsFields(fields...)
	}
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	layout := GetLogLayout(ctx)
	if layout != nil {
		layout.Error(msg, fields...)
	}
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	layout := GetLogLayout(ctx)
	if layout != nil {
		layout.Warn(msg, fields...)
	}
}
