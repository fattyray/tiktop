package views

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"tiktop/db"
	"tiktop/midware"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//前面db这个定义成了那个包括初始化所有db的文件的包了，所有这里链接数据库不能再用db了，连着写糊涂了好几次
	dsn := "wcr123:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	dbx, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": "failed to connect database"})
		return
	}

	user := db.User{}
	result := dbx.Where("Username = ?", username).First(&user)
	//没有找到就进行注册
	if result.Error == gorm.ErrRecordNotFound {
		//用户的名字就用它提供的，密码直接明文存储，如果后续有需求，可以改成sha256之类的
		//id就使用雪花算法生成(保证唯一性)
		node, err := snowflake.NewNode(1)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "failed to generate snowflake"})
		}
		idx := uint(node.Generate().Int64())
		Newuser := db.User{Userid: idx, Username: username, Password: password}
		if err := dbx.Create(&Newuser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": 2, "status_msg": "Failed to create user"})
			return
		}
		tokenx, err := midware.CreateToken(string(idx), password)
		c.JSON(http.StatusOK, gin.H{"status_code": 0, "status_msg": "Account registered successfully", "user_id": idx, "token": tokenx})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"StatusCode": 3, "Message": "Account already exists"})
		return
	}

}
