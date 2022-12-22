package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type AuthMiddleware struct{}

func (a *AuthMiddleware) MiddlewareHandlerFunc(option *defs.ControllerOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(option.AuthCodes) > 0 {
			token, err := uctx.GetJwtClaimsByBearerToken(c)
			if err != nil {
				response.FailWithMessage(c, err)
				c.Abort()
				return
			}
			if len(token.TenantID) > 0 {
				if token.TenantID != uctx.GetTenantID(c) {
					response.FailWithMessage(c, errs.PermissionDeniedErr)
					c.Abort()
					return
				}
				auths := token.Auths
				var hasAuth = false
				for _, v := range auths {
					for _, authCode := range option.AuthCodes {
						if v == authCode {
							hasAuth = true
							break
						}
					}
				}
				if !hasAuth {
					response.FailWithMessage(c, errs.PermissionDeniedErr)
					c.Abort()
					return
				}
			}
		}
		c.Next()
	}
}
