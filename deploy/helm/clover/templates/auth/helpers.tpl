{{- define "auth.service.fullname" -}}
{{- if .Values.auth.fullnameOverride -}}
{{- .Values.auth.fullnameOverride | trunc 63 | trimSuffix "-" | quote -}}
{{- else -}}
{{- print .Release.Name "-auth" | quote -}}
{{- end }}
{{- end }}

{{- define "auth.server.config" }}
# 服务配置
svcConf:
  # 服务模式,1正式2测试
  svcMode: 2
  # 服务部署名
  svcName: "{{- print .Release.Name "-auth" -}}"
  # 服务端口号
  svcPort: {{ .Values.auth.service.targetPort }}
  # 服务实例序号
  svcNum: 1
  # 服务版本号
  svcVersion: "v0.0.1"

# consul配置
consulConf:
  # 服务配置中心路径
  consulAddr: "consul:8500"
  # 注册时限
  registerTTL: 2
  # 注册间隔
  registerInterval: 1

# 日志配置
logConf:
  # 日志级别
  level: "debug"
  # 输出文件夹
  director: "/filebeat/logs"
  # 最大天数
  maxAge: 7
  # 是否输出到控制台
  logInConsole: 1
{{ end }}

{{- define "auth.sentinel.config" }}
# 是否开启熔断限流，1是2否
enabled: 1

# 服务全局熔断限流策略
svcRuleStrategy:
  # 限流配置
  currentLimitRules:
    # 计算策略，0阈值控制1预热计算
    - tokenCalculateStrategy: 0
      # 控制策略，0拒绝1排队
      controlBehavior: 1
      # 阈值
      threshold: 1000
      # 排队最大等待时间
      maxQueueingTimeMs: 1000
      # 统计周期
      statIntervalInMs: 1000
  # 热点限流配置
  hotspotRules:
    # 流控指标类型，0并发数，1请求数
    - metricType: 0
      # 控制策略，0拒绝1排队
      controlBehavior: 1
      # 参数键
      paramKey: TenantId
      # 阈值
      threshold: 10
      # 排队最大等待时间
      maxQueueingTimeMs: 1000
{{ end }}

{{- define "auth.global.config" }}
# redis配置
redisConfig:
  # 是否启用redis，1是2否
  enabled: 1
  # redis域名地址
  addr: "redis:6379"
  # redis密码
  password: "clover"
  # redis子库
  db: 1
# jwt配置
jwtConfig:
  # 签名密钥
  signingKey: "clover"
  # 过期时限，单位秒
  expiresTime: 14400
# 遥测体系配置
otelConfig:
  # 是否启用遥测体系，1是2否
  enabled: 1
  # collector的GRPC端口
  collectorGrpcEndpoint: "otel-collector:4317"
nsqConfig:
  # 是否开启NSQ消息队列，1是2否
  enabled: 1
  # 用于生产者使用的nsqd服务地址
  producerAddr: "nsqd:4150"
  # 用于消费者使用的nsqlookupd服务地址
  consumerAddr: "nsqlookupd:4161"
{{ end }}

{{- define "auth.db.config" }}
postgres:
  path: postgres-master
  port: 5432
  config: sslmode=disable TimeZone=Asia/Shanghai
  dbName: clover_auth
  username: root
  password: 123456
  maxIdleConns: 10
  maxOpenConns: 50
{{ end }}