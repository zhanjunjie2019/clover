package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
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
// @Router /user-create [post]
func (u *UserCreateController) Handle(c *gin.Context) {
	var userCreateReqVO vo.UserCreateReqVO
	ctx, err := uctx.ShouldBindJSON(c, &userCreateReqVO)
	if err == nil {
		var id defs.ID
		id, err = u.UserApp.UserCreate(ctx, userCreateReqVO.UserName, userCreateReqVO.Password)
		if err == nil {
			response.SuccWithDetailed(c, vo.UserCreateRspVO{
				UserId: id.UInt64(),
			})
			return
		}
	}
	response.FailWithMessage(c, err)
}
