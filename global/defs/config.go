package defs

// IConfigDefine 动态配置定义
type IConfigDefine interface {
	// GetOption 配置标识
	GetOption() ConfigOption
	// ReloadConfig 配置变更通知，传入指针对象
	ReloadConfig(any) error
	// Unmarshal 反序列化对象，返回指针对象
	Unmarshal([]byte) (any, error)
}

type ConfigOption struct {
	// ConfigKey 配置标识
	ConfigKey string
	// CanLoadByConsul 是否从配置中心读取
	CanLoadByConsul bool
	// ConfigFileName 配置文件在根路径下的路径
	ConfigFileName string
}
