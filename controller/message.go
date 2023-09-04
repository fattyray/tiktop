package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tiktop/entity"
	"tiktop/service"
	"tiktop/util"
)

func MessageChat(c *gin.Context) {
	//校验token并获取当前用户id
	token := c.Query("token")
	claims, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	fromUserId, _ := strconv.Atoi(claims.UserId)
	//获取目标用户id
	toUserId, err1 := strconv.Atoi(c.Query("to_user_id"))

	msgTime, err2 := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "toUserId解析失败"})
		return
	}

	fmt.Printf("fromUserId=%v\n", int64(fromUserId))
	list, err := service.ChatList(int64(fromUserId), int64(toUserId), msgTime)

	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "获取聊天记录失败"})
		return
	}

	c.JSON(http.StatusOK, entity.ChatResponse{
		Response:    entity.Response{StatusCode: 0},
		MessageList: list,
	})
}

func MessageAction(c *gin.Context) {

	//校验token并获取当前用户id
	token := c.Query("token")
	claims, err := util.Gettoken(token)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "token error"})
		return
	}
	fromUserId, _ := strconv.Atoi(claims.UserId)
	//获取目标用户id
	toUserId, err1 := strconv.Atoi(c.Query("to_user_id"))

	content := c.Query("content")
	actionType, err2 := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "userId或actiontype解析失败"})
		return
	}
	if content == "" || actionType != 1 {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "参数错误"})
		return
	}
	message, err := service.AddMessage(int64(fromUserId), int64(toUserId), content)
	if err != nil {
		log.Println(message, err)
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: message + err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: message})
}
