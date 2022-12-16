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

type RoleCreateController struct {
	RoleApp app.RoleAppIOCInterface `singleton:""`
}

func (r *RoleCreateController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/role-create",
		HttpMethod:   http.MethodPost,
		AuthCode:     consts.AdminAuth,
	}
}

// Handle 创建角色，需要租户管理员权限
// @Tags role
// @Summary 创建角色，需要租户管理员权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param Tenant-ID header string true "租户ID"
// @Param data body vo.RoleCreateReqVO true "创建角色"
// @Success 200 {object} response.Response{data=vo.RoleCreateRspVO}
// @Router /auth/role-create [post]
func (r *RoleCreateController) Handle(c *gin.Context) {
	var (
		command cmd.RoleCreateCmd
		result  cmd.RoleCreateResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.RoleCreateReqVO{}, nil, &command)
	if err == nil {
		result, err = r.RoleApp.RoleCreate(ctx, command)
		if err == nil {
			var rspVO vo.RoleCreateRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
			return
		}
	}
	response.FailWithMessage(c, err)
}
