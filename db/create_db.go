package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktop/sql_dsn"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type User struct {
	Userid   uint
	Username string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`

	gorm.Model
}
type Video struct {
	Videoid  uint
	Videourl string `gorm:"type:varchar(500)"`
	Coverurl string `gorm:"type:varchar(500)"`
	Title    string `gorm:"type:varchar(100)"`
	Userid   uint
	gorm.Model
}
type Comment struct {
	Commentid uint
	Userid    uint
	Videoid   uint
	Content   string `gorm:"type:varchar(300)"`
	gorm.Model
}
type Follow struct {
	Userid       uint
	FollowUserid uint
	gorm.Model
}

type Like struct {
	Userid  uint
	Videoid uint
	gorm.Model
}

func Init_db() error {
	// 连接数据库
	// 去 sql_dsn 里面修改 电脑数据库
	//dsn := "root:123456@(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := sql_dsn.GetDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//自动迁移
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Video{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Follow{})
	db.AutoMigrate(&Like{})
	fmt.Println("db init")
	//u1 := User{Id: 1, Name: "张三", Gender: "男", Hobby: "学习"}
	//db.Create(&u1) //创建
	return nil
}
