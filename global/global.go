package global

import (
	"gorm.io/gorm"
)

var ( // 全局变量
	DB *gorm.DB // 数据库接口
)
