package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktop/entity"
	"tiktop/service"
	"tiktop/util"
	"time"
)

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	claims, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	currentId, _ := strconv.Atoi(claims.UserId)

	videoId, _ := strconv.Atoi(c.Query("video_id"))
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	if actionType == 1 {
		commentText := c.Query("comment_text")
		err = service.AddComment(int64(currentId), int64(videoId), commentText)
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 2, StatusMsg: "add comment failed"})
			return
		}
		var userdata entity.UserData
		userdata, err = service.UserInfoByUserId(int64(currentId))
		userdata.IsFollow, err = service.QueryFollowOrNot(int64(currentId), int64(currentId))
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 3, StatusMsg: "follow data query failed"})
			return
		}
		c.JSON(http.StatusOK, entity.CommentActionResponse{
			Response: entity.Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			CommentResponse: entity.CommentResponse{
				Id:         int64(currentId),
				User:       userdata,
				Content:    commentText,
				CreateDate: time.Now().Format("01-02"),
			},
		})
	} else if actionType == 2 {
		commentId, _ := strconv.Atoi(c.Query("comment_id"))
		err = service.CancelComment(int64(currentId), int64(videoId), int64(commentId))
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 4, StatusMsg: "cancel comment failed"})
			return
		}
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "cancel comment success"})
	}
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	claims, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	currentId, _ := strconv.Atoi(claims.UserId)
	comments, err := service.GetCommentListByVideoId(int64(currentId), int64(videoId))
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 2, StatusMsg: "get comment list failed"})
		return
	}
	c.JSON(http.StatusOK, entity.CommentListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		Comments: comments,
	})
}
