package controller

import (
	"github.com/gin-gonic/gin"
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

type TenantCreateController struct {
	TenantApp app.TenantAppIOCInterface `singleton:""`
}

func (t *TenantCreateController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/tenant-create",
		HttpMethod:   http.MethodPost,
	}
}

// Handle 创建租户
// @Tags tenant
// @Summary 创建租户
// @accept application/json
// @Produce application/json
// @Param data body vo.TenantCreateReqVO true "创建租户"
// @Success 200 {object} response.Response{data=vo.TenantCreateRspVO}
// @Router /tenant-create [post]
func (t *TenantCreateController) Handle(c *gin.Context) {
	var tenantCreateReqVO vo.TenantCreateReqVO
	ctx, err := uctx.ShouldBindJSON(c, &tenantCreateReqVO)
	if err == nil {
		var tid, secretKey string
		if tenantCreateReqVO.AccessTokenTimeLimit == 0 {
			tenantCreateReqVO.AccessTokenTimeLimit = 7200
		}
		if tenantCreateReqVO.RefreshTokenTimeLimit == 0 {
			tenantCreateReqVO.RefreshTokenTimeLimit = 86400
		}
		tid, secretKey, err = t.TenantApp.TenantCreate(ctx,
			tenantCreateReqVO.TenantID, tenantCreateReqVO.TenantName, tenantCreateReqVO.RedirectUrl,
			tenantCreateReqVO.AccessTokenTimeLimit, tenantCreateReqVO.RefreshTokenTimeLimit,
		)
		if err == nil {
			response.SuccWithDetailed(c, vo.TenantCreateRspVO{
				TenantID:  tid,
				SecretKey: secretKey,
			})
			return
		}
	}
	response.FailWithMessage(c, err)
}
