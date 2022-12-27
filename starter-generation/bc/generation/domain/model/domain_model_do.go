package model

import "github.com/zhanjunjie2019/clover/global/defs"

func NewModule(id defs.ID, value ModuleValue) Module {
	return &module{
		id:    id,
		value: value,
	}
}

type Module interface {
	ID() defs.ID
	// FullValue 核心数据
	FullValue() ModuleValue
}

type module struct {
	id    defs.ID
	value ModuleValue
}

func (m module) ID() defs.ID {
	return m.id
}

func (m module) FullValue() ModuleValue {
	return m.value
}

type ModuleValue struct {
	// 生成代码保存路径
	RootPackagePath string
	// 模块名
	ModuleName string
	// 服务默认端口
	ServerPort uint16
	// 是否加载配置中心配置，1=是，其他=否
	EnabledConfigByConsul uint8
	// 是否启用遥测链路追踪，1=是，其他=否Z
	EnabledOpenTelemetry uint8
}
