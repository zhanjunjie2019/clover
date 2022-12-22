package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
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

type RolePermissionAssignmentController struct {
	RoleApp app.RoleAppIOCInterface `singleton:""`
}

func (r *RolePermissionAssignmentController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/role-permission-assignment",
		HttpMethod:   http.MethodPost,
		AuthCodes:    []string{consts.AdminAuth, consts.SAdminAuth},
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
// @Router /auth/role-permission-assignment [post]
func (r *RolePermissionAssignmentController) Handle(c *gin.Context) {
	var (
		command cmd.RolePermissionAssignmentCmd
		result  cmd.RolePermissionAssignmentResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.RolePermissionAssignmentReqVO{},
		func(arg *vo.RolePermissionAssignmentReqVO) bool {
			arg.AuthCodes = lo.Uniq(arg.AuthCodes)
			return true
		}, &command)
	if err == nil {
		result, err = r.RoleApp.RolePermissionAssignment(ctx, command)
		if err == nil {
			var rspVO vo.RolePermissionAssignmentRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
