package confs

import "github.com/zhanjunjie2019/clover/global/utils"

type DBType string

const (
	PostgreSQL = "postgres"
	MySQL      = "mysql"
)

type DBConfig struct {
	// Enabled 是否启用redis，1是2否
	Enabled Enabled `yaml:"enabled"`
	// DBType
	DbType DBType `yaml:"dbType"`
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

// AllAaccordance 完全一致
func (d DBConfig) AllAaccordance(d2 DBConfig) bool {
	if !d.MainAaccordance(d2) ||
		d.MaxIdleConns != d2.MaxIdleConns ||
		d.MaxOpenConns != d2.MaxOpenConns {
		return false
	}
	return true
}

// MainAaccordance 主要配置一致
func (d DBConfig) MainAaccordance(d2 DBConfig) bool {
	if !utils.AllAaccordance(
		string(d.DbType), string(d2.DbType),
		d.Path, d2.Path,
		d.Port, d2.Port,
		d.Config, d2.Config,
		d.Dbname, d2.Dbname,
		d.Username, d2.Username,
		d.Password, d2.Password,
	) {
		return false
	}
	return true
}
