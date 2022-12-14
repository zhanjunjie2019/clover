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

type SadminTokenCreateController struct {
	SadminApp app.SadminAppIOCInterface `singleton:""`
}

func (s *SadminTokenCreateController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/sadmin-token-create",
		HttpMethod:   http.MethodPost,
	}
}

// Handle 获得超管Token
// @Tags tenant
// @Summary 获得超管Token
// @accept application/json
// @Produce application/json
// @Param data body vo.SadminTokenCreateReqVO true "获得超管Token"
// @Success 200 {object} response.Response{data=vo.SadminTokenCreateRspVO}
// @Router /sadmin-token-create [post]
func (s *SadminTokenCreateController) Handle(c *gin.Context) {
	var reqVO vo.SadminTokenCreateReqVO
	ctx, err := uctx.ShouldBindJSON(c, &reqVO)
	if err == nil {
		var (
			accessToken               string
			accessTokenExpirationTime int64
		)
		accessToken, accessTokenExpirationTime, err = s.SadminApp.SadminTokenCreate(ctx, reqVO.SecretKey)
		if err == nil {
			response.SuccWithDetailed(c, vo.SadminTokenCreateRspVO{
				AccessToken:               accessToken,
				AccessTokenExpirationTime: accessTokenExpirationTime,
			})
			return
		}
	}
	response.FailWithMessage(c, err)
}
