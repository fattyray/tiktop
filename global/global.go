package global

import (
	"gorm.io/gorm"
)

var ( // 全局变量
	DB          *gorm.DB // 数据库接口
	MaxNumVideo = 30
	PATH_VIDEO  = "./publish/video/"
	PATH_COVER  = "./publish/cover/" // 封面相对路径
)
