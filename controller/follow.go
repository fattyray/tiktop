package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktop/entity"
	"tiktop/service"
	"tiktop/util"
)

func RelationAction(c *gin.Context) {
	//校验token并获取当前用户id
	token := c.Query("token")
	claims, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	currentId, _ := strconv.Atoi(claims.UserId)
	//获取目标用户id
	uid, _ := strconv.Atoi(c.Query("to_user_id"))
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	err = service.DoOrCancelFollow(int64(currentId), int64(uid), int32(actionType))
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 2, StatusMsg: "follow action failed"})
		return
	}
	c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "action success"})
}

func FollowList(c *gin.Context) {
	//校验token并获取当前用户id
	token := c.Query("token")
	claims, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	currentId, _ := strconv.Atoi(claims.UserId)
	//获取目标用户id
	uid, _ := strconv.Atoi(c.Query("user_id"))
	if currentId != uid {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "no permission"})
		return
	}
	followUserList, err := service.GetFollowUserListByUserId(int64(uid))
	//封装返回
	c.JSON(http.StatusOK, entity.FollowListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: followUserList,
	})
}

func FollowerList(c *gin.Context) {
	//校验token并获取当前用户id
	token := c.Query("token")
	claims, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	currentId, _ := strconv.Atoi(claims.UserId)
	//获取目标用户id
	uid, _ := strconv.Atoi(c.Query("user_id"))
	if currentId != uid {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "no permission"})
		return
	}
	followUserList, err := service.GetFollowerListByUserId(int64(uid))
	//封装返回
	c.JSON(http.StatusOK, entity.FollowListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: followUserList,
	})
}

// 好友列表是关注登录用户的粉丝而已
func FriendList(c *gin.Context) {
	FollowerList(c)
}
