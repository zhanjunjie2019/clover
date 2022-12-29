package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/adapter/controller/vo"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/app"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/app/cmd"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/auth/infr/bcconsts"
	"net/http"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type TenantTokenCreateController struct {
	TenantApp app.TenantAppIOCInterface `singleton:""`
}

func (t *TenantTokenCreateController) GetOption() defs.ControllerOptions {
	return defs.NewControllerOptions(
		defs.RelativePath(bcconsts.ModuleCode+"/tenant-token-create"),
		defs.HttpMethod(http.MethodPost),
	)
}

// Handle 创建租户管理员Token
// @Tags tenant
// @Summary 创建租户Token
// @accept application/json
// @Produce application/json
// @Param data body vo.TenantTokenCreateReqVO true "创建租户Token"
// @Success 200 {object} response.Response{data=vo.TenantTokenCreateRspVO}
// @Router /auth/tenant-token-create [post]
func (t *TenantTokenCreateController) Handle(c *gin.Context) {
	var (
		command cmd.TenantTokenCreateCmd
		result  cmd.TenantTokenCreateResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.TenantTokenCreateReqVO{}, func(arg *vo.TenantTokenCreateReqVO) bool {
		// 当前时间
		seconds := time.Now().Unix()
		// token过期时间，要么为0，要么大于当前时间
		if arg.AccessTokenExpirationTime > 0 && seconds > arg.AccessTokenExpirationTime {
			return false
		}
		return true
	}, &command)
	if err == nil {
		result, err = t.TenantApp.TenantTokenCreate(ctx, command)
		if err == nil {
			var rspVO vo.TenantTokenCreateRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
