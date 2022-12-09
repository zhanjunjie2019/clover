package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
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

type RolePermissionAssignmentController struct {
	RoleApp app.RoleAppIOCInterface `singleton:""`
}

func (r *RolePermissionAssignmentController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/role-permission-assignment",
		HttpMethod:   http.MethodPost,
		AuthCode:     consts.AdminAuth,
	}
}

// Handle 角色赋予资源许可，需要租户管理员权限
// @Tags role
// @Summary 角色赋予资源许可，需要租户管理员权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param Tenant-ID header string true "租户ID"
// @Param data body vo.RolePermissionAssignmentReqVO true "角色赋予资源许可"
// @Success 200 {object} response.Response{data=vo.RolePermissionAssignmentRspVO}
// @Router /role-permission-assignment [post]
func (r *RolePermissionAssignmentController) Handle(c *gin.Context) {
	var reqVO vo.RolePermissionAssignmentReqVO
	ctx, err := uctx.ShouldBindJSON(c, &reqVO)
	if err == nil {
		reqVO.AuthCodes = lo.Uniq(reqVO.AuthCodes)
		var id = defs.ID(reqVO.RoleID)
		id, err = r.RoleApp.RolePermissionAssignment(ctx, id, reqVO.RoleCode, reqVO.AuthCodes)
		if err == nil {
			response.SuccWithDetailed(c, vo.RolePermissionAssignmentRspVO{
				RoleID: id.UInt64(),
			})
			return
		}
	}
	response.FailWithMessage(c, err)
}
