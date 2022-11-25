package configs

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-auth/bc/infr/bcconsts"
	"gopkg.in/yaml.v3"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IConfigDefine

type AuthConfigDefine struct{}

func (a *AuthConfigDefine) GetOption() defs.ConfigOption {
	return defs.ConfigOption{
		ConfigKey:       bcconsts.AuthConfigDefineKey,
		CanLoadByConsul: true,
		ConfigFileName:  bcconsts.AuthConfigFileName,
	}
}

func (a *AuthConfigDefine) ReloadConfig(config any) error {
	ac := config.(*AuthConfig)
	SetAuthConfig(*ac)
	return nil
}

func (a *AuthConfigDefine) Unmarshal(data []byte) (any, error) {
	var ac AuthConfig
	err := yaml.Unmarshal(data, &ac)
	return &ac, err
}

// --以下是缓存的配置模型--

var authConfig AuthConfig

func GetAuthConfig() AuthConfig {
	return authConfig
}

func SetAuthConfig(ac AuthConfig) {
	authConfig = ac
}

// --以下是定义的配置模型--

type AuthConfig struct {
	SuperAdmin *AuthSuperAdminConfig `yaml:"superAdmin"`
}

type AuthSuperAdminConfig struct {
	// TenantID 租户ID
	TenantID string `yaml:"tenantID"`
	// SecretKey 租户密钥
	SecretKey string `yaml:"secretKey"`
	// AccessTokenTimeLimit 访问Token有效时限
	AccessTokenTimeLimit uint64 `yaml:"accessTokenTimeLimit"`
}
