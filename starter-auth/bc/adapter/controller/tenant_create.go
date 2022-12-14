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

type TenantCreateController struct {
	TenantApp app.TenantAppIOCInterface `singleton:""`
}

func (t *TenantCreateController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/tenant-create",
		HttpMethod:   http.MethodPost,
		AuthCode:     consts.SAdminAuth,
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
// @Router /tenant-create [post]
func (t *TenantCreateController) Handle(c *gin.Context) {
	var reqVO vo.TenantCreateReqVO
	ctx, err := uctx.ShouldBindJSON(c, &reqVO)
	if err == nil {
		var tid, secretKey string
		if reqVO.AccessTokenTimeLimit == 0 {
			reqVO.AccessTokenTimeLimit = 7200
		}
		tid, secretKey, err = t.TenantApp.TenantCreate(ctx,
			reqVO.TenantID, reqVO.TenantName, reqVO.RedirectUrl,
			reqVO.AccessTokenTimeLimit,
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
