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

type TenantCreateController struct {
	TenantApp app.TenantAppIOCInterface `singleton:""`
}

func (t *TenantCreateController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/tenant-create",
		HttpMethod:   http.MethodPost,
		AuthCodes:    []string{consts.SAdminAuth},
	}
}

// Handle 创建租户
// @Tags tenant
// @Summary 创建租户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body vo.TenantCreateReqVO true "创建租户"
// @Success 200 {object} response.Response{data=vo.TenantCreateRspVO}
// @Router /auth/tenant-create [post]
func (t *TenantCreateController) Handle(c *gin.Context) {
	var (
		command cmd.TenantCreateCmd
		result  cmd.TenantCreateResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.TenantCreateReqVO{}, nil, &command)
	if err == nil {
		result, err = t.TenantApp.TenantCreate(ctx, command)
		if err == nil {
			var rspVO vo.TenantCreateRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
