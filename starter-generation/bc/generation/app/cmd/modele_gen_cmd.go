package cmd

type ModuleGenCmd struct {
	// 生成代码保存路径
	RootPackagePath string `json:"rootPackagePath"`
	// 模块名，必须满足英文字母{2}~{20}位
	ModuleName string `json:"moduleName"`
	// 服务默认端口，0~65535
	ServerPort uint16 `json:"serverPort"`
	// 是否加载配置中心配置，1=是，其他=否
	EnabledConfigByConsul uint8 `json:"enabledConfigByConsul"`
	// 是否启用遥测链路追踪，1=是，其他=否
	EnabledOpenTelemetry uint8 `json:"enabledOpenTelemetry"`
	// 领域限界上下文列表
	BoundedContexts []DomainBoundedContext `json:"boundedContexts"`
}

type DomainBoundedContext struct {
	// 限界名称，长度必须满足{2}~{30}位
	BoundedContextName string `json:"entityName"`
	// 领域聚合
	Aggregations []DomainAggregation `json:"aggregations"`
}

type DomainAggregation struct {
	// 聚合根
	AggregationRoot DomainEntity `json:"aggregationRoot"`
}

type DomainEntity struct {
	// 实体名，长度必须满足{2}~{30}位
	EntityName string `json:"entityName"`
	// 实体
	Entitys []DomainEntity `json:"entitys"`
	// 值对象
	ValueObjects []DomainValueObject `json:"valueObjects"`
}

type DomainValueObject struct {
	// 值对象名，长度必须满足{2}~{30}位
	ValObjName string `json:"valObjName"`
	// 值对象链接类型，1=一对一，其他=一对多
	ValRelType uint8 `json:"valRelType"`
	// 值基础类型，0=非基础值对象，1=string，2=uint8，3=uint16，4=uint32，5=uint64
	ValObjBaseType uint8 `json:"valObjBaseType"`
	// 值对象
	ValueObjects []DomainValueObject `json:"valueObjects"`
}

type ModuleGenResult struct {
}
