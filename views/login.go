package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"tiktop/db"
	"tiktop/midware"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	dsn := "wcr123:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	dbx, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dbx.AutoMigrate(&db.User{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": "failed to connect database"})
		return
	}
	user := db.User{}
	result := dbx.Where("Username = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		//没有该账号
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "Account not found"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": "User doesn't exist error"})
		return
	} else if password == user.Password {
		fmt.Println(int(user.Userid))
		tokenx, _ := midware.CreateToken(string(user.Userid), string(password))
		c.JSON(http.StatusOK, gin.H{"status_code": 0, "status_msg": "Login successful", "user_id": user.Userid, "token": tokenx})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status_code": 2, "status_msg": "Incorrect password"})
	}
}
