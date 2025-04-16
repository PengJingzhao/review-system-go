package config

// Mysql 是一个结构体，用于定义 MySQL 数据库的配置项
type Mysql struct {
	Host         string `yaml:"host"`         // 数据库主机地址
	Port         string `yaml:"port"`         // 数据库端口号
	Config       string `yaml:"config"`       // 数据库连接的额外配置参数
	DB           string `yaml:"db"`           // 数据库名称
	Username     string `yaml:"username"`     // 数据库用户名
	Password     string `yaml:"password"`     // 数据库密码
	LogLevel     string `yaml:"log_level"`    // 日志级别
	MaxIdleConns int    `yaml:"maxIdleConns"` // 设置空闲连接池中连接的最大数量
	MaxOpenConns int    `yaml:"maxOpenConns"` // 设置数据库最大打开连接数
}

// Dsn 方法生成 MySQL 的 Data Source Name (DSN)
// DSN 是用于连接数据库的字符串
func (m *Mysql) Dsn() string {
	// 拼接 DSN 格式：用户名:密码@tcp(主机:端口)/数据库名称?附加配置
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DB + "?" + m.Config
}
