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

type UserRoleAssignmentController struct {
	UserApp app.UserAppIOCInterface `singleton:""`
}

func (u *UserRoleAssignmentController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/user-role-assignment",
		HttpMethod:   http.MethodPost,
		AuthCodes:    []string{consts.AdminAuth, consts.SAdminAuth},
	}
}

// Handle 用户赋予角色，需要租户管理员权限
// @Tags user
// @Summary 用户赋予角色，需要租户管理员权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param Tenant-ID header string true "租户ID"
// @Param data body vo.UserRoleAssignmentReqVO true "用户赋予角色"
// @Success 200 {object} response.Response{data=vo.UserRoleAssignmentRspVO}
// @Router /auth/user-role-assignment [post]
func (u *UserRoleAssignmentController) Handle(c *gin.Context) {
	var (
		command cmd.UserRoleAssignmentCmd
		result  cmd.UserRoleAssignmentResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.UserRoleAssignmentReqVO{},
		func(arg *vo.UserRoleAssignmentReqVO) bool {
			arg.RoleCodes = lo.Uniq(arg.RoleCodes)
			return true
		}, &command)
	if err == nil {
		result, err = u.UserApp.UserRoleAssignment(ctx, command)
		if err == nil {
			var rspVO vo.UserRoleAssignmentRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
