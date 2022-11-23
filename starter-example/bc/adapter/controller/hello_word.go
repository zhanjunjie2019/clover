package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/starter-example/bc/adapter/controller/vo"
	"github.com/zhanjunjie2019/clover/starter-example/bc/app"
	"github.com/zhanjunjie2019/clover/starter-example/bc/infr/bcconsts"
	"net/http"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type HelloWordController struct {
	ExampleApp app.ExampleAppIOCInterface `singleton:""`
}

func (h *HelloWordController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/hellow-word",
		HttpMethod:   http.MethodGet,
	}
}

// Handle 接口描述
// @Tags example
// @Summary 接口描述
// @accept application/json
// @Produce application/json
// @Param data query vo.HelloWordReqVO true "参数描述"
// @Success 200 {object} response.Response{data=vo.HelloWordRspVO}
// @Router /hellow-word [get]
func (h *HelloWordController) Handle(c *gin.Context) {
	var helloWordReqVO vo.HelloWordReqVO
	ctx, err := uctx.ShouldBindQuery(c, &helloWordReqVO)
	if err == nil {
		var greetings string
		greetings, err = h.ExampleApp.ExampleHellowWord(ctx, helloWordReqVO.FirstName, helloWordReqVO.LastName)
		if err == nil {
			response.SuccWithDetailed(c, vo.HelloWordRspVO{
				Greetings: greetings,
			})
			return
		}
	}
	response.FailWithMessage(c, err)
}
