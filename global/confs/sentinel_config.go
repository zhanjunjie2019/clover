package confs

import (
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/core/hotspot"
)

var sentinelConfig SentinelConfig

func GetSentinelConfig() SentinelConfig {
	return sentinelConfig
}

func SetSentinelConfig(sc SentinelConfig) {
	sentinelConfig = sc
}

type SentinelConfig struct {
	// Enabled 是否开启熔断限流,1是2否
	Enabled Enabled `yaml:"enabled"`
	// SvcRuleStrategy 服务全局熔断限流策略
	SvcRuleStrategy RuleStrategy `yaml:"svcRuleStrategy"`
	// ApisRuleStrategys 接口级熔断限流策略
	ApisRuleStrategys map[string]RuleStrategy `yaml:"apisRuleStrategys"`
}

type RuleStrategy struct {
	// FlowRules限流配置
	FlowRules []FlowRule `yaml:"currentLimitRules"`
	// HotspotRules 热点限流配置
	HotspotRules []HotspotRule `yaml:"hotspotRules"`
}

// FlowRule 无参限流配置
type FlowRule struct {
	// TokenCalculateStrategy 当前流量控制器的Token计算策略。Direct表示直接使用字段 Threshold 作为阈值；WarmUp表示使用预热方式计算Token的阈值。
	TokenCalculateStrategy flow.TokenCalculateStrategy `yaml:"tokenCalculateStrategy"`
	// ControlBehavior 表示流量控制器的控制策略；Reject表示超过阈值直接拒绝，Throttling表示匀速排队。
	ControlBehavior flow.ControlBehavior `yaml:"controlBehavior"`
	// Threshold 表示流控阈值；如果字段 StatIntervalInMs 是1000(也就是1秒)，那么Threshold就表示QPS，流量控制器也就会依据资源的QPS来做流控。
	Threshold float64 `yaml:"threshold"`
	// MaxQueueingTimeMs 匀速排队的最大等待时间，该字段仅仅对 Throttling ControlBehavior 生效；
	MaxQueueingTimeMs uint32 `yaml:"maxQueueingTimeMs"`
	// WarmUpPeriodSec 预热的时间长度，该字段仅仅对 WarmUp 的 TokenCalculateStrategy 生效；
	WarmUpPeriodSec uint32 `yaml:"warmUpPeriodSec"`
	// WarmUpColdFactor 预热的因子，默认是3，该值的设置会影响预热的速度，该字段仅仅对 WarmUp 的 TokenCalculateStrategy 生效；
	WarmUpColdFactor uint32 `yaml:"warmUpColdFactor"`
	// StatIntervalInMs 规则对应的流量控制器的独立统计结构的统计周期。如果 StatIntervalInMs 是1000，也就是统计QPS。
	StatIntervalInMs uint32 `yaml:"statIntervalInMs"`
}

// HotspotRule 热点参数限流配置
type HotspotRule struct {
	// MetricType 流控指标类型，支持两种：请求数和并发数
	MetricType hotspot.MetricType `yaml:"metricType"`
	// ControlBehavior 流控的效果，仅在请求数模式下有效。支持两种：快速失败和匀速+排队模式
	ControlBehavior hotspot.ControlBehavior `yaml:"controlBehavior"`
	// ParamIndex 热点参数的索引，对应 WithArgs(args ...any) 中的参数索引位置，从 0 开始
	ParamIndex int `yaml:"paramIndex"`
	// ParamKey 参数键
	ParamKey string `yaml:"paramKey"`
	// Threshold 限流阈值（针对每个热点参数）
	Threshold int64 `yaml:"threshold"`
	// MaxQueueingTimeMs 最大排队等待时长（仅在匀速排队模式 + QPS 下生效）
	MaxQueueingTimeMs int64 `yaml:"maxQueueingTimeMs"`
	// BurstCount 静默值(仅在快速失败模式 + QPS 下生效)
	BurstCount int64 `yaml:"burstCount"`
	// DurationInSec 统计结构填充新的 token 的时间间隔 (仅在请求数(QPS)流控模式下生效)
	DurationInSec int64 `yaml:"durationInSec"`
	// ParamsMaxCapacity 统计结构的容量最大值（Top N）
	ParamsMaxCapacity int64 `yaml:"paramsMaxCapacity"`
	// SpecificItems 特定参数的特殊阈值配置，可以针对指定的参数值单独设置限流阈值，不受前面 Threshold 阈值的限制。
	SpecificItems map[any]int64 `yaml:"specificItems"`
}
