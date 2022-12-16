package configs

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/bcconsts"
	"gopkg.in/yaml.v3"
)

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:implements=github.com/zhanjunjie2019/clover/global/defs.IConfigDefine

type ExampleConfigDefine struct{}

func (e *ExampleConfigDefine) GetOption() defs.ConfigOption {
	return defs.ConfigOption{
		ConfigKey:       bcconsts.ExampleConfigDefineKey,
		CanLoadByConsul: false,
		ConfigFileName:  bcconsts.ExampleConfigFileName,
	}
}

func (e *ExampleConfigDefine) ReloadConfig(config any) error {
	ec := config.(*ExampleConfig)
	SetExampleConfig(*ec)
	return nil
}

func (e *ExampleConfigDefine) Unmarshal(data []byte) (any, error) {
	var ec ExampleConfig
	err := yaml.Unmarshal(data, &ec)
	return &ec, err
}

// --以下是缓存的配置模型--

var exampleConfig ExampleConfig

func GetExampleConfig() ExampleConfig {
	return exampleConfig
}

func SetExampleConfig(ec ExampleConfig) {
	exampleConfig = ec
}

// --以下是定义的配置模型--

type ExampleConfig struct {
	Aa *ExampleAaConfig `yaml:"aa"`
}

type ExampleAaConfig struct {
	Bb string `yaml:"bb"`
}
