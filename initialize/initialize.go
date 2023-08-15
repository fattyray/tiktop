package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktop/entity"
	"tiktop/global"
)

func InitDB() error {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/tiktop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//自动迁移
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Video{})
	db.AutoMigrate(&entity.Comment{})
	db.AutoMigrate(&entity.Follow{})
	db.AutoMigrate(&entity.Like{})
	fmt.Println("db init")
	//u1 := User{Id: 1, Name: "张三", Gender: "男", Hobby: "学习"}
	//db.Create(&u1) //创建
	global.DB = db
	return nil
}
