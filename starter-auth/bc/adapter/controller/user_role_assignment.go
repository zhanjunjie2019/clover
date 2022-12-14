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

type UserRoleAssignmentController struct {
	UserApp app.UserAppIOCInterface `singleton:""`
}

func (u *UserRoleAssignmentController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/user-role-assignment",
		HttpMethod:   http.MethodPost,
		AuthCode:     consts.AdminAuth,
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
// @Router /user-role-assignment [post]
func (u *UserRoleAssignmentController) Handle(c *gin.Context) {
	var reqVO vo.UserRoleAssignmentReqVO
	ctx, err := uctx.ShouldBindJSON(c, &reqVO)
	if err == nil {
		var id = defs.ID(reqVO.UserID)
		id, err = u.UserApp.UserRoleAssignment(ctx, id, reqVO.UserName, reqVO.RoleCodes)
		if err == nil {
			response.SuccWithDetailed(c, vo.UserRoleAssignmentRspVO{
				UserID: id.UInt64(),
			})
			return
		}
	}
	response.FailWithMessage(c, err)
}
