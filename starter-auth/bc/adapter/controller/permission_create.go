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

type PermissionCreateController struct {
	PermissionApp app.PermissionAppIOCInterface `singleton:""`
}

func (p *PermissionCreateController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/permission-create",
		HttpMethod:   http.MethodPost,
		AuthCode:     consts.SAdminAuth,
	}
}

// Handle 创建资源权限许可
// @Tags permission
// @Summary 创建资源权限许可
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param Tenant-ID header string true "租户ID"
// @Param data body vo.PermissionReqVO true "创建资源权限许可"
// @Success 200 {object} response.Response{data=vo.PermissionRspVO}
// @Router /permission-create [post]
func (p *PermissionCreateController) Handle(c *gin.Context) {
	var reqVO vo.PermissionReqVO
	ctx, err := uctx.ShouldBindJSON(c, &reqVO)
	if err == nil {
		var id defs.ID
		id, err = p.PermissionApp.PermissionCreate(ctx, reqVO.PermissionName, reqVO.AuthCode)
		if err == nil {
			response.SuccWithDetailed(c, vo.PermissionRspVO{
				PermissionID: id.UInt64(),
			})
			return
		}
	}
	response.FailWithMessage(c, err)
}
