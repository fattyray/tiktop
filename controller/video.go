package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktop/entity"
	"tiktop/service"
	"tiktop/util"
)

type videoListResponse struct {
	entity.Response
	VideoList []entity.VideoResponse `json:"video_list"`
}

func PublishList(c *gin.Context) {
	//校验token并获取当前用户id
	token := c.Query("token")
	_, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	uid, _ := strconv.Atoi(c.Query("user_id"))
	videos, err := service.GetPostVideoListByUserId(int64(uid))
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 2, StatusMsg: "get liked video list failed"})
		return
	}
	//封装返回
	c.JSON(http.StatusOK, videoListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videos,
	})
}
