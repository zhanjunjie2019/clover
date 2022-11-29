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
// @Router /role-create [post]
func (r *RoleCreateController) Handle(c *gin.Context) {
	var roleCreateReqVO vo.RoleCreateReqVO
	ctx, err := uctx.ShouldBindJSON(c, &roleCreateReqVO)
	if err == nil {
		var id defs.ID
		id, err = r.RoleApp.RoleCreate(ctx, roleCreateReqVO.RoleName, roleCreateReqVO.RoleCode)
		if err == nil {
			response.SuccWithDetailed(c, vo.RoleCreateRspVO{
				RoleId: id.UInt64(),
			})
			return
		}
	}
	response.FailWithMessage(c, err)
}
