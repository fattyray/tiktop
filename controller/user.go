package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktop/entity"
	"tiktop/service"
	"tiktop/util"
)

type UserLoginResponse struct {
	entity.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	entity.Response
	User entity.UserData `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := service.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "register failed"})
		return
	}
	tokenx, _ := util.CreateToken(strconv.FormatUint(uint64(user.UserId), 10), password)
	c.JSON(http.StatusOK,
		UserLoginResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "Account registered successfully"},
			UserId:   user.UserId,
			Token:    tokenx,
		})
	return
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// 从数据库查询用户信息
	user, err := service.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK,
			UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "用户名或密码错误"},
			})
		return
	}
	// 生成对应 token
	tokenx, errlog := util.CreateToken(strconv.Itoa(int(user.UserId)), string(password))
	if errlog != nil {
		c.JSON(http.StatusOK,
			UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: errlog.Error()},
			})
		return
	}
	c.Set("currentId", user.UserId)
	// 返回成功并生成响应 json
	c.JSON(http.StatusOK,
		UserLoginResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "Login successfully"},
			UserId:   user.UserId,
			Token:    tokenx,
		})
}

func UserInfo(c *gin.Context) {
	value := c.Query("user_id")
	//println(value)
	userId, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "invalid id format"})
		return
	}
	token := c.Query("token")
	claims, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	var userdata entity.UserData
	userdata, err = service.UserInfoByUserId(int64(userId))
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 2, StatusMsg: "user data query failed"})
		return
	}
	currentId, _ := strconv.Atoi(claims.UserId)
	isFollow, err := service.QueryFollowOrNot(int64(currentId), int64(userId))
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 2, StatusMsg: "follow data query failed"})
		return
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "OK"},
		User: entity.UserData{
			userdata.UserId,
			userdata.Name,
			userdata.FollowCount,
			userdata.FollowerCount,
			isFollow,
			userdata.Avatar,
			userdata.BackgroundImage,
			userdata.Signature,
			userdata.TotalFavorited,
			userdata.WorkCount,
			userdata.FavoriteCount,
		},
	})
}
