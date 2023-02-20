package confs

var globalConfig GlobalConfig

func GetGlobalConfig() GlobalConfig {
	return globalConfig
}

func SetGlobalConfig(gc GlobalConfig) {
	globalConfig = gc
}

type GlobalConfig struct {
	// RedisConfig redis配置
	RedisConfig RedisConfig `yaml:"redisConfig"`
	// JwtConfig jwt配置
	JwtConfig JwtConfig `yaml:"jwtConfig"`
	// OtelConfig 遥测体系配置
	OtelConfig OtelConfig `yaml:"otelConfig"`
	// NsqConfig NSQ消息队列监听
	NsqConfig NsqConfig `yaml:"nsqConfig"`
}

// RedisConfig redis配置
type RedisConfig struct {
	// Enabled 是否启用redis，1是2否
	Enabled Enabled `env:"REDIS_ENABLED" yaml:"enabled"`
	// Addr redis域名地址
	Addr string `env:"REDIS_ADDR" yaml:"addr"`
	// Password redis密码
	Password string `env:"REDIS_POSSWORD" yaml:"password"`
	// DB redis子库
	DB uint8 `env:"REDIS_DB" yaml:"db"`
}

// JwtConfig jwt配置
type JwtConfig struct {
	// SigningKey 签名密钥
	SigningKey string `env:"JWT_SIGNING_KEY" yaml:"signingKey"`
	// ExpiresTime 过期时限，单位秒
	ExpiresTime uint64 `env:"JWT_EXPIRES_TIME" yaml:"expiresTime"`
}

// OtelConfig 遥测体系配置
type OtelConfig struct {
	// Enabled 是否启用遥测体系，1是2否
	Enabled Enabled `env:"OTEL_ENABLED" yaml:"enabled"`
	// CollectorGrpcEndpoint collector端口
	CollectorGrpcEndpoint string `env:"OTEL_COLLECTOR_ENDPOINT" yaml:"collectorGrpcEndpoint"`
}

// NsqConfig NSQ消息队列配置
type NsqConfig struct {
	// Enabled 是否开启NSQ消息队列，1是2否
	Enabled Enabled `env:"NSQ_ENABLED" yaml:"enabled"`
	// ProducerAddr 用于生产者使用的nsqd服务地址
	ProducerAddr string `env:"NSQ_NSQD_ADDR" yaml:"producerAddr"`
	// ConsumerAddr 用于消费者使用的nsqlookupd服务地址
	ConsumerAddr string `env:"NSQ_LOOKUPD_ADDR" yaml:"consumerAddr"`
}
