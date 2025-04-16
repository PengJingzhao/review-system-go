package config

// Logger 是一个结构体，用于定义日志配置项
type Logger struct {
	Level        string `yaml:"level"`        // 日志级别，例如 info、debug、error 等
	Prefix       string `yaml:"prefix"`       // 日志前缀，用于标识日志的来源或类型
	Director     string `yaml:"director"`     // 日志文件目录
	ShowLine     bool   `yaml:"show-line"`    // 是否显示日志所在的代码行信息
	LogInConsole bool   `yaml:"logInConsole"` // 是否在控制台输出日志
}
