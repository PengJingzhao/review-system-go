package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"review-system-go/config"
)

var (
	CONFIG *config.Config // 配置文件
	DB     *gorm.DB       // // 数据库连接对象
	LOG    *logrus.Logger // 日志
)
