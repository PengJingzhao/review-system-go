package config

import "fmt"

// Server 是一个结构体，用于定义服务器相关的配置项
type Server struct {
	Host   string `yaml:"host"`   // 服务器主机地址，例如 "0.0.0.0" 或 "127.0.0.1"
	Port   string `yaml:"port"`   // 服务器端口号，例如 "8080"
	Env    string `yaml:"env"`    // 运行环境，例如 "release"、"debug" 或 "test"
	JwtKey string `yaml:"jwtKey"` // 用于 JWT 的密钥
}

// GetAddr 返回服务器的完整地址
func (s *Server) GetAddr() string {
	// 使用 fmt.Sprintf 拼接 "主机:端口" 的格式
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
