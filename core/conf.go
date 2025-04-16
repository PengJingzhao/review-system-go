package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"review-system-go/config"
	"review-system-go/global"
)

// InitConf 读取settings.yaml文件
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}

	// 读取 YAML 文件内容
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("读取 YAML 配置文件失败: %s", err))
	}

	// 解析 YAML 文件到 Config 结构体
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("结构体映射错误: %s", err)
	}

	// 打印日志或标记加载成功
	log.Println("配置文件加载成功!")
	global.CONFIG = c // 设置全局变量
}
