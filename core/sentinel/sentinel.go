package sentinel

import (
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/core/hotspot"
	"github.com/samber/lo"
	"github.com/zhanjunjie2019/clover/global/confs"
	"sync"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
// +ioc:autowire:constructFunc=InitSentinelLoader

type SentinelLoader struct {
	// flowRules 无参限流
	flowRules map[string]*flow.Rule
	// hotspotRoles 热点限流
	hotspotRoles map[string]*hotspot.Rule
	sync.Mutex
}

func InitSentinelLoader(s *SentinelLoader) (*SentinelLoader, error) {
	s.CleanBufferRules()
	return s, nil
}

// CleanBufferRules 清空缓存规则
func (s *SentinelLoader) CleanBufferRules() {
	s.Lock()
	defer s.Unlock()
	s.flowRules = make(map[string]*flow.Rule, 0)
	s.hotspotRoles = make(map[string]*hotspot.Rule, 0)
}

// AppendServerRules 添加服务级限流缓存
func (s *SentinelLoader) AppendServerRules() {
	s.Lock()
	defer s.Unlock()
	strategy := confs.GetSentinelConfig().SvcRuleStrategy
	svcName := confs.GetServerConfig().SvcConf.SvcName
	s.appendRules(svcName, strategy)
}

// AppendApiRules 添加接口级限流缓存
func (s *SentinelLoader) AppendApiRules(sentinelStrategy string) {
	s.Lock()
	defer s.Unlock()
	apisRuleStrategys := confs.GetSentinelConfig().ApisRuleStrategys
	strategy := apisRuleStrategys[sentinelStrategy]
	svcName := confs.GetServerConfig().SvcConf.SvcName
	s.appendRules(svcName+"-"+sentinelStrategy, strategy)
}

// appendRules 加载规则
func (s *SentinelLoader) appendRules(resource string, ruleStrategy confs.RuleStrategy) {
	if len(ruleStrategy.FlowRules) > 0 {
		for _, v := range ruleStrategy.FlowRules {
			rule := &flow.Rule{
				Resource:               resource,
				TokenCalculateStrategy: v.TokenCalculateStrategy,
				ControlBehavior:        v.ControlBehavior,
				Threshold:              v.Threshold,
				MaxQueueingTimeMs:      v.MaxQueueingTimeMs,
				WarmUpPeriodSec:        v.WarmUpPeriodSec,
				WarmUpColdFactor:       v.WarmUpColdFactor,
				StatIntervalInMs:       v.StatIntervalInMs,
			}
			s.flowRules[resource] = rule
		}
	}
	if len(ruleStrategy.HotspotRules) > 0 {
		for _, v := range ruleStrategy.HotspotRules {
			rule := &hotspot.Rule{
				Resource:          resource,
				MetricType:        v.MetricType,
				ControlBehavior:   v.ControlBehavior,
				ParamIndex:        v.ParamIndex,
				ParamKey:          v.ParamKey,
				Threshold:         v.Threshold,
				MaxQueueingTimeMs: v.MaxQueueingTimeMs,
				BurstCount:        v.BurstCount,
				DurationInSec:     v.DurationInSec,
				ParamsMaxCapacity: v.ParamsMaxCapacity,
				SpecificItems:     v.SpecificItems,
			}
			s.hotspotRoles[resource] = rule
		}
	}
}

// LoadSentinelRules 加载全部规则
func (s *SentinelLoader) LoadSentinelRules() error {
	s.Lock()
	defer s.Unlock()
	if len(s.flowRules) > 0 {
		_, err := flow.LoadRules(lo.Values(s.flowRules))
		return err
	}
	if len(s.hotspotRoles) > 0 {
		_, err := hotspot.LoadRules(lo.Values(s.hotspotRoles))
		return err
	}
	return nil
}
