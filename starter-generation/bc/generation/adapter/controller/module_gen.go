package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/response"
	"github.com/zhanjunjie2019/clover/global/uctx"
	"github.com/zhanjunjie2019/clover/global/utils"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/adapter/controller/vo"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/app"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/app/cmd"
	"github.com/zhanjunjie2019/clover/starter-generation/bc/generation/infr/bcconsts"
	"net/http"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type ModuleGenController struct {
	GenerationApp app.GenerationAppIOCInterface `singleton:""`
}

func (m *ModuleGenController) GetOption() defs.ControllerOptions {
	return defs.NewControllerOptions(
		defs.RelativePath(bcconsts.ModuleCode+"/module-gen"),
		defs.HttpMethod(http.MethodPost),
		defs.AuthCodes(consts.SAdminAuth),
	)
}

// Handle 生成完整服务代码
// @Tags gen
// @Summary 生成完整服务代码
// @accept application/json
// @Produce application/json
// @Param data query vo.ModuleGenReqVO true "生成完整服务代码"
// @Success 200 {object} response.Response{data=vo.ModuleGenRspVO}
// @Router /gen/module-gen [get]
func (m *ModuleGenController) Handle(c *gin.Context) {
	var (
		command cmd.ModuleGenCmd
		result  cmd.ModuleGenResult
	)
	ctx, err := uctx.ShouldBindJSON(c, &vo.ModuleGenReqVO{}, nil, &command)
	if err == nil {
		result, err = m.GenerationApp.ModuleGen(ctx, command)
		if err == nil {
			var rspVO vo.ModuleGenRspVO
			err = utils.CopyObject(&result, &rspVO)
			if err == nil {
				response.SuccWithDetailed(c, rspVO)
				return
			}
		}
	}
	response.FailWithMessage(c, err)
}
