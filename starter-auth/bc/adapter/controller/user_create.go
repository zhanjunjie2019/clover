package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/app"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/bcconsts"
	"net/http"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IController

type UserCreate struct {
	UserApp app.UserAppIOCInterface `singleton:""`
}

func (u *UserCreate) GetOption() defs.ControllerOption {
	return defs.ControllerOption{
		RelativePath: bcconsts.ModuleCode + "/user-create",
		HttpMethod:   http.MethodPost,
		AuthCode:     consts.AdminAuth,
	}
}

// Handle 创建用户，需要租户管理员权限
func (u *UserCreate) Handle(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
