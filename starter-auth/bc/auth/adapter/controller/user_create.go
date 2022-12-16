package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
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

type UserCreateController struct {
	UserApp app.UserAppIOCInterface `singleton:""`
}

func (u *UserCreateController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/user-create",
		HttpMethod:   http.MethodPost,
		AuthCode:     consts.AdminAuth,
	}
}

// Handle 创建用户，需要租户管理员权限
// @Tags user
// @Summary 创建用户，需要租户管理员权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param Tenant-ID header string true "租户ID"
// @Param data body vo.UserCreateReqVO true "创建用户"
// @Success 200 {object} response.Response{data=vo.UserCreateRspVO}
// @Router /auth/user-create [post]
func (u *UserCreateController) Handle(c *gin.Context) {
	var (
		command cmd.UserCreateCmd
		result  cmd.UserCreateResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.UserCreateReqVO{}, nil, &command)
	if err == nil {
		result, err = u.UserApp.UserCreate(ctx, command)
		if err == nil {
			var rspVO vo.UserCreateRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
