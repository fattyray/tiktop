package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktop/entity"
	"tiktop/service"
	"tiktop/util"
)

// 点赞和取消赞操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	claims, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	currentId, _ := strconv.Atoi(claims.UserId)
	err = service.GiveOrCancelLike(int64(currentId), int64(videoId), int32(actionType))
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 2, StatusMsg: "like action failed"})
		return
	}
	c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "action success"})
}

func FavoriteList(c *gin.Context) {
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
	videos, err := service.GetLikeVideoListByUserId(int64(uid), int64(currentId))
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 2, StatusMsg: "get liked video list failed"})
		return
	}
	//封装返回
	c.JSON(http.StatusOK, entity.FavoriteListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videos,
	})
}
