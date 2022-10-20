package confs

var dbConfig DBConfig

func GetDBConfig() DBConfig {
	return dbConfig
}

func SetDBConfig(dc DBConfig) {
	dbConfig = dc
}

type DBConfig struct {
	// pgsql的数据库配置
	Postgres *GeneralDBConf `yaml:"postgres"`
	// mysql的数据库配置
	Mysql *GeneralDBConf `yaml:"mysql"`
}

type GeneralDBConf struct {
	// 服务器地址
	Path string `yaml:"path"`
	//端口
	Port string `yaml:"port"`
	// 高级配置
	Config string `yaml:"config"`
	// 数据库名
	Dbname string `yaml:"dbName"`
	// 数据库用户名
	Username string `yaml:"username"`
	// 数据库密码
	Password string `yaml:"password"`
	// 空闲中的最大连接数
	MaxIdleConns int `yaml:"maxIdleConns"`
	// 打开到数据库的最大连接数
	MaxOpenConns int `yaml:"maxOpenConns"`
}
