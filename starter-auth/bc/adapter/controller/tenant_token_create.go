package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/adapter/controller/vo"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/app"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/bcconsts"
	"net/http"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type TenantTokenCreateController struct {
	TenantApp app.TenantAppIOCInterface `singleton:""`
}

func (t *TenantTokenCreateController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/tenant-token-create",
		HttpMethod:   http.MethodPost,
	}
}

// Handle 创建租户Token
// @Tags tenant
// @Summary 创建租户Token
// @accept application/json
// @Produce application/json
// @Param data body vo.TenantTokenCreateReqVO true "创建租户Token"
// @Success 200 {object} response.Response{data=vo.TenantTokenCreateRspVO}
// @Router /tenant-token-create [post]
func (t *TenantTokenCreateController) Handle(c *gin.Context) {
	var tenantTokenCreateReqVO vo.TenantTokenCreateReqVO
	ctx, err := uctx.ShouldBindJSON(c, &tenantTokenCreateReqVO)
	if err == nil {
		// 参数验证flag
		pass := true
		// 当前时间
		seconds := time.Now().Unix()
		// token过期时间，要么为0，要么大于当前时间
		if tenantTokenCreateReqVO.AccessTokenExpirationTime > 0 && seconds > tenantTokenCreateReqVO.AccessTokenExpirationTime {
			err = errs.ReqParamsErr
			pass = false
		}
		if pass {
			var (
				accessToken               string
				accessTokenExpirationTime int64
			)
			accessToken, accessTokenExpirationTime, err = t.TenantApp.TenantTokenCreate(ctx, tenantTokenCreateReqVO.TenantID, tenantTokenCreateReqVO.SecretKey, tenantTokenCreateReqVO.AccessTokenExpirationTime)
			if err == nil {
				response.SuccWithDetailed(c, vo.TenantTokenCreateRspVO{
					AccessToken:               accessToken,
					AccessTokenExpirationTime: accessTokenExpirationTime,
				})
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
