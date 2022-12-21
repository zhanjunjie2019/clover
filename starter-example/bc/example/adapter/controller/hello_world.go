package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/adapter/controller/vo"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/app"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/app/cmd"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/bcconsts"
	"net/http"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type HelloWorldController struct {
	ExampleApp app.ExampleAppIOCInterface `singleton:""`
}

func (h *HelloWorldController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/hellow-world",
		HttpMethod:   http.MethodGet,
	}
}

// Handle 接口描述
// @Tags example
// @Summary 接口描述
// @accept application/json
// @Produce application/json
// @Param data query vo.HelloWorldReqVO true "参数描述"
// @Success 200 {object} response.Response{data=vo.HelloWorldRspVO}
// @Router /example/hellow-world [get]
func (h *HelloWorldController) Handle(c *gin.Context) {
	var (
		command cmd.HelloWorldCmd
		result  cmd.HelloWorldResult
	)
	ctx, err := uctx.ShouldBindQuery(c, &vo.HelloWorldReqVO{}, nil, &command)
	if err == nil {
		result, err = h.ExampleApp.ExampleHellowWorld(ctx, command)
		if err == nil {
			var rspVO vo.HelloWorldRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
