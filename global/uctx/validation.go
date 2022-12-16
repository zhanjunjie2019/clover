package uctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/utils"
)

var validate = validator.New()

// ShouldBindJSON 解析并验证请求body参数，返回链路span上下文
// check:额外自定义的VO校验方法
func ShouldBindJSON[T any](c *gin.Context, obj T, check func(arg T) bool, rs ...any) (ctx context.Context, err error) {
	err = c.ShouldBindJSON(obj)
	if err != nil {
		return c, errs.ReqParamsErr
	}
	err = validate.Struct(obj)
	if err != nil {
		return c, errs.ReqParamsErr
	}
	if check != nil {
		if !check(obj) {
			return c, errs.ReqParamsErr
		}
	}
	if len(rs) > 0 {
		var last any = obj
		for i := range rs {
			err = utils.CopyObject(last, rs[i])
			if err != nil {
				return c, errs.NewUnknownErr(err)
			}
			last = rs[i]
		}
	}
	ctx = GetSpanContext(c)
	return
}

// ShouldBindQuery 解析并验证请求query参数，返回链路span上下文
// check:额外自定义的VO校验方法
func ShouldBindQuery[T any](c *gin.Context, obj T, check func(arg T) bool, rs ...any) (ctx context.Context, err error) {
	err = c.ShouldBindQuery(obj)
	if err != nil {
		return c, errs.ReqParamsErr
	}
	err = validate.Struct(obj)
	if err != nil {
		return c, errs.ReqParamsErr
	}
	if check != nil {
		if !check(obj) {
			return c, errs.ReqParamsErr
		}
	}
	if len(rs) > 0 {
		var last any = obj
		for i := range rs {
			err = utils.CopyObject(last, rs[i])
			if err != nil {
				return c, errs.NewUnknownErr(err)
			}
			last = rs[i]
		}
	}
	ctx = GetSpanContext(c)
	return
}
