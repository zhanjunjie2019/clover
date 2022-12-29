package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/adapter/controller/vo"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/app"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/app/cmd"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/bcconsts"
	"net/http"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type UserTokenByAuthcodeController struct {
	UserApp app.UserAppIOCInterface `singleton:""`
}

func (u *UserTokenByAuthcodeController) GetOption() defs.ControllerOptions {
	return defs.NewControllerOptions(
		defs.RelativePath(bcconsts.ModuleCode+"/user-token-by-authcode"),
		defs.HttpMethod(http.MethodPost),
	)
}

// Handle 使用授权码获取用户token信息
// @Tags user
// @Summary 使用授权码获取用户token信息
// @accept application/json
// @Produce application/json
// @Param Tenant-ID header string true "租户ID"
// @Param data body vo.UserTokenByAuthcodeReqVO true "使用授权码获取用户token信息"
// @Success 200 {object} response.Response{data=vo.UserTokenByAuthcodeRspVO}
// @Router /auth/user-token-by-authcode [post]
func (u *UserTokenByAuthcodeController) Handle(c *gin.Context) {
	var (
		command cmd.UserTokenByAuthcodeCmd
		result  cmd.UserTokenByAuthcodeResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.UserTokenByAuthcodeReqVO{}, nil, &command)
	if err == nil {
		result, err = u.UserApp.UserTokenByAuthcode(ctx, command)
		if err == nil {
			var rspVO vo.UserTokenByAuthcodeRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
