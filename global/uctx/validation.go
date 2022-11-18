package uctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zhanjunjie2019/clover/global/errs"
)

var validate = validator.New()

// ShouldBindJSON 解析并验证请求body参数，返回链路span上下文
func ShouldBindJSON(c *gin.Context, obj any) (ctx context.Context, err error) {
	err = c.ShouldBindJSON(obj)
	if err != nil {
		return c, errs.ReqParamsErr
	}
	err = validate.Struct(obj)
	if err != nil {
		return c, errs.ReqParamsErr
	}
	ctx = GetSpanContext(c)
	return
}
