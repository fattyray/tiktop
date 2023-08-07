package views

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"tiktop/db"
	"tiktop/midware"
)

// 其中加了omitempty的部分的是空的话就不加入，后面五个内容这部分内容目前是没有的，你们做了社交相关内容后进行完善

func UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")
	_, err := midware.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 3, "status_msg": "jwt error"})
		return
	}

	dsn := "wcr123:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	dbx, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": "failed to connect database"})
		return
	}
	user := db.User{}

	result := dbx.Where("Userid = ?", userid).First(&user)
	//resultx := dbx.Where("Userid = ?", token_id).First(&jwt_user)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": -2, "status_msg": "no such user"})
		return
	}
	type userdata struct {
		Id            int    `json:"id,omitempty"`
		Name          string `json:"name,omitempty"`
		FollowCount   int64  `json:"follow_count,omitempty"`
		FollowerCount int64  `json:"follower_count,omitempty"`
		IsFollow      bool   `json:"is_follow,omitempty"`
		LikedCount    int64  `json:"total_favorited,omitempty"`
		LikeCount     int64  `json:"favorite_count,omitempty"`
	}
	tmpid, _ := strconv.Atoi(userid)
	myinfo := userdata{Id: tmpid, Name: user.Username}

	jsonData, err1 := json.Marshal(myinfo)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": -5, "status_msg": "json.Marshal error"})
	}
	fmt.Println(string(jsonData))
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "status_msg": "OK", "user": string(jsonData)})
	return

}
