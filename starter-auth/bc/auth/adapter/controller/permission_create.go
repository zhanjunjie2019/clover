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

type PermissionCreateController struct {
	PermissionApp app.PermissionAppIOCInterface `singleton:""`
}

func (p *PermissionCreateController) GetOption() defs.ControllerOptions {
	return defs.NewControllerOptions(
		defs.RelativePath(bcconsts.ModuleCode+"/permission-create"),
		defs.HttpMethod(http.MethodPost),
		defs.AuthCodes(consts.SAdminAuth),
	)
}

// Handle 创建资源权限许可
// @Tags permission
// @Summary 创建资源权限许可
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body vo.PermissionReqVO true "创建资源权限许可"
// @Success 200 {object} response.Response{data=vo.PermissionRspVO}
// @Router /auth/permission-create [post]
func (p *PermissionCreateController) Handle(c *gin.Context) {
	var (
		command cmd.PermissionCreateCmd
		result  cmd.PermissionCreateResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.PermissionReqVO{}, nil, &command)
	if err == nil {
		result, err = p.PermissionApp.PermissionCreate(ctx, command)
		if err == nil {
			var rspVO vo.PermissionRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
