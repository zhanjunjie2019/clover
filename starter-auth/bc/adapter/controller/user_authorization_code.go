package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/adapter/controller/vo"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/app"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/bcconsts"
	"net/http"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type UserAuthorizationCodeController struct {
	UserApp app.UserAppIOCInterface `singleton:""`
}

func (u *UserAuthorizationCodeController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/user-authorization-code",
		HttpMethod:   http.MethodPost,
	}
}

// Handle 用户登录获得授权码，注意授权码不是Token，不能直接用于访问接口
// @Tags auth
// @Summary 登录验证用户账号密码，验证通过后在Redis保存一个授权码60秒有效，关联用户信息。用以可以用授权码接口换取登录Token。
// @accept application/json
// @Produce application/json
// @Param "Tenant-ID" header string true "租户ID"
// @Param data body vo.UserAuthorizationCodeReqVO true "用户登录获得授权码，注意授权码不是Token，不能直接用于访问接口"
// @Success 200 {object} response.Response{data=vo.UserAuthorizationCodeRspVO}
// @Router /user-authorization-code [post]
func (u *UserAuthorizationCodeController) Handle(c *gin.Context) {
	var userAuthorizationCodeReqVO vo.UserAuthorizationCodeReqVO
	ctx, err := uctx.ShouldBindJSON(c, &userAuthorizationCodeReqVO)
	if err == nil {
		code, url, err := u.UserApp.UserAuthorizationCode(ctx,
			userAuthorizationCodeReqVO.UserName,
			userAuthorizationCodeReqVO.Password,
			userAuthorizationCodeReqVO.RedirectUrl)
		if err == nil {
			response.SuccWithDetailed(c, vo.UserAuthorizationCodeRspVO{
				AuthorizationCode: code,
				RedirectUrl:       url,
			})
			return
		}
	}
	response.FailWithMessage(c, err)
}
