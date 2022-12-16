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
// @Router /auth/sadmin-token-create [post]
func (s *SadminTokenCreateController) Handle(c *gin.Context) {
	var (
		command cmd.SadminTokenCreateCmd
		result  cmd.SadminTokenCreateResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.SadminTokenCreateReqVO{}, nil, &command)
	if err == nil {
		result, err = s.SadminApp.SadminTokenCreate(ctx, command)
		if err == nil {
			var rspVO vo.SadminTokenCreateRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
