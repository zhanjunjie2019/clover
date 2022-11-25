package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/adapter/controller/vo"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/bcconsts"
	"net/http"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type PermissionCreateController struct {
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
	var permissionReqVO vo.PermissionReqVO
	_, err := uctx.ShouldBindJSON(c, &permissionReqVO)
	if err == nil {

	}
	response.FailWithMessage(c, err)
}
