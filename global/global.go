package global

import (
	"gorm.io/gorm"
)

var ( // 全局变量
	DB          *gorm.DB            // 数据库接口
	MaxNumVideo = 30                // 一次最大搜查视频量
	PATH_VIDEO  = "./public/video/" // 视频保存相对路径
	PATH_COVER  = "./public/cover/" // 封面相对路径
	HEAD_URL    = "http://"
	VIDEO_URL   = "/static/video/"
	COVER_URL   = "/static/cover/"
)
