package uctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
)

func WithValueTenantID(ctx context.Context, tenantID string) context.Context {
	return context.WithValue(ctx, consts.CtxTenantIDVar, tenantID)
}

func GetTenantID(ctx context.Context) string {
	c, ok := ctx.(*gin.Context)
	if ok {
		value, ok := c.Get(consts.CtxTenantIDVar)
		if ok {
			s, ok := value.(string)
			if ok {
				return s
			}
		}
		tenantID := c.GetHeader(consts.TenantIDHeaderKey)
		if len(tenantID) == 0 {
			tenantID = c.Param("TenantId")
		}
		if len(tenantID) > 0 {
			c.Set(consts.CtxTenantIDVar, tenantID)
		}
		return tenantID
	} else {
		value := ctx.Value(consts.CtxTenantIDVar)
		if value != nil {
			s, ok := value.(string)
			if ok {
				return s
			}
		}
	}
	return ""
}
