package response

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"go.uber.org/zap"
	"net/http"
)

type Response struct {
	// 错误码：成功[0],底层的错误[1000-1999],平台层级的错误[2000-2999]
	Code int `json:"code"`
	// 数据
	Data any `json:"data,omitempty"`
	// 详细信息
	Msg string `json:"msg,omitempty"`
}

func Result(c *gin.Context, code int, data any, msg string) {
	jbs, _ := json.Marshal(Response{
		code,
		data,
		msg,
	})
	rs := string(jbs)
	uctx.AppendLogsFields(c, zap.String("rspBody", rs))
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, rs)
}

func SuccWithDetailed(c *gin.Context, data any) {
	Result(c, 0, data, "")
}

func FailWithMessage(c *gin.Context, err error) {
	FailWithDetailed(c, err, nil)
}

func FailWithDetailed(c *gin.Context, err error, data any) {
	uctx.Warn(c, err.Error(), zap.Error(err))
	if unifiedErr, ok := err.(*errs.UnifiedError); ok {
		Result(c, unifiedErr.Code(), data, unifiedErr.ShowError())
	} else {
		Result(c, errs.UnknownErrorCode, data, fmt.Sprintf("[%d]%s", errs.UnknownErrorCode, err.Error()))
	}
}
