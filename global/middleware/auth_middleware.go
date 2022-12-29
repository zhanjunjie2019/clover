package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"go-micro.dev/v4/server"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type AuthMiddleware struct{}

// MiddlewareWrapHandlerByAuthCodes 有需要的gRPC接口，要手动在注册时加上
func (a *AuthMiddleware) MiddlewareWrapHandlerByAuthCodes(authCodes []string) server.HandlerWrapper {
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx, token, err := uctx.GetJwtClaimsByGrpcCtx(ctx)
			if err != nil {
				return err
			}
			ctx, tenantID := uctx.GetTenantIDByGrpcCtx(ctx)
			err = a.filter(token, authCodes, tenantID)
			if err != nil {
				return err
			}
			return handlerFunc(ctx, req, rsp)
		}
	}
}

func (a *AuthMiddleware) MiddlewareHandlerFunc(option *defs.ControllerOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := uctx.GetJwtClaimsByBearerToken(c)
		if err != nil {
			response.FailWithMessage(c, err)
			c.Abort()
			return
		}
		err = a.filter(token, option.AuthCodes, uctx.GetTenantID(c))
		if err != nil {
			response.FailWithMessage(c, err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func (a *AuthMiddleware) filter(token defs.JwtClaims, authCodes []string, tenantID string) error {
	if len(token.TenantID) > 0 {
		if token.TenantID != tenantID {
			return errs.PermissionDeniedErr
		}
		auths := token.Auths
		var hasAuth = false
		for _, v := range auths {
			for _, authCode := range authCodes {
				if v == authCode {
					hasAuth = true
					break
				}
			}
		}
		if !hasAuth {
			return errs.PermissionDeniedErr
		}
	}
	return nil
}
