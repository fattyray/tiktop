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
	"tiktop/sql_dsn"
)

// 其中加了omitempty的部分的是空的话就不加入，后面五个内容这部分内容目前是没有的，你们做了社交相关内容后进行完善

func UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	fmt.Println(userid)
	token := c.Query("token")
	_, err := midware.Gettoken(token)
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{"status_code": 3, "status_msg": "null"})
		c.JSON(http.StatusOK, db.Response{
			StatusCode: 3,
			StatusMsg:  "null",
		})
		return
	}

	dsn := sql_dsn.GetDsn()
	//dsn := "root:123456@(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	dbx, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": "null"})
		c.JSON(http.StatusOK, db.Response{
			StatusCode: -1,
			StatusMsg:  "null",
		})
		return
	}
	user := db.User{}

	result := dbx.Where("Userid = ?", userid).First(&user)
	//resultx := dbx.Where("Userid = ?", token_id).First(&jwt_user)
	if result.Error != nil {
		//c.JSON(http.StatusOK, gin.H{"status_code": -2, "status_msg": "null"})
		c.JSON(http.StatusOK, db.Response{
			StatusCode: -2,
			StatusMsg:  "null",
		})
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
		//c.JSON(http.StatusOK, gin.H{"status_code": -5, "status_msg": "null"})
		c.JSON(http.StatusOK, db.Response{
			StatusCode: -5,
			StatusMsg:  "null",
		})
	}
	fmt.Println(string(jsonData))
	//这个是要等所以的社交信息填好之后才传入user
	//c.JSON(http.StatusOK, gin.H{"status_code": 0, "status_msg": "string", "user": string(jsonData)})
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "status_msg": "string"})
	//c.JSON(http.StatusOK, db.Response{
	//	StatusCode: 0,
	//	StatusMsg:  "string",
	//})
	return

}
