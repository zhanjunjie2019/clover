package confs

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

var serverConfig ServerConfig

func GetServerConfig() ServerConfig {
	return serverConfig
}

func SetServerConfig(sc ServerConfig) {
	serverConfig = sc
}

type ServerConfig struct {
	// SvcConf 服务配置
	SvcConf SvcConf `yaml:"svcConf"`
	// ConsulConf consul配置
	ConsulConf ConsulConf `yaml:"consulConf"`
	// LogConf 日志配置
	LogConf LogConf `yaml:"logConf"`
}

type SvcMode uint8

const (
	Release SvcMode = 1 + iota
	Debug
)

type SvcConf struct {
	// SvcMode 服务模式,1正式2测试
	SvcMode SvcMode `env:"SVC_MODE" yaml:"svcMode"`
	// SvcName 服务部署名
	SvcName string `env:"SVC_NAME" yaml:"svcName"`
	// SvcNum 服务实例序号
	SvcNum uint8 `env:"SVC_NUM" yaml:"svcNum"`
	// SvcPort 服务端口号
	SvcPort uint16 `env:"SVC_PORT" yaml:"svcPort"`
	// SvcVersion 服务版本号
	SvcVersion string `env:"SVC_VERSION" yaml:"svcVersion"`
}

type ConsulConf struct {
	// ConsulAddr 服务配置中心路径
	ConsulAddr string `env:"CONSUL_ADDR" yaml:"consulAddr"`
	// RegisterTTL 注册时限
	RegisterTTL uint16 `env:"CONSUL_REGISTER_TTL" yaml:"registerTTL"`
	// RegisterInterval 注册间隔
	RegisterInterval uint16 `env:"CONSUL_REGISTER_INTERVAL" yaml:"registerInterval"`
}

type LogConf struct {
	// Level 日志级别
	Level string `env:"LOG_LEVEL" yaml:"level"`
	// Director 输出文件夹
	Director string `env:"LOG_DIRECTOR" yaml:"director"`
	// MaxAge 最大天数
	MaxAge uint8 `env:"LOG_MAX_AGE" yaml:"maxAge"`
	// LogInConsole 是否输出到控制台
	LogInConsole uint8 `env:"LOG_IN_CONSOLE" yaml:"logInConsole"`
}

func (z *LogConf) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
