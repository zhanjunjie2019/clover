package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/bcconsts"
	"net/http"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type PermissionCreateController struct {
}

func (p *PermissionCreateController) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/permission-create",
		HttpMethod:   http.MethodPost,
		AuthCode:     consts.SAdminAuth,
	}
}

func (p *PermissionCreateController) Handle(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
