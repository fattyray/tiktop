package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktop/db"
	"tiktop/global"
	"tiktop/service"
	"time"
)

type feedResponse struct {
	Response db.Response
	//StatusCode int        `json:"status_code,omitempty"`
	VideoList []db.Video `json:"video_list,omitempty"`
	NextTime  int64      `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	//获取 last_time, 找不到时使用当前时间
	LastTimeStr := c.DefaultQuery("last_time", "")
	var LastTime int64
	CurrentTime := time.Now().Unix()
	if LastTimeStr != "" {
		LastTimeTemp, err := time.Parse(time.RFC3339, LastTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status_code": -1, "status_msg": "Fail to get the last time."})
			return
		}
		LastTime = LastTimeTemp.Unix()
	} else {
		LastTime = CurrentTime
	}

	// 判断此时的视频列表是否为空
	var videoList []db.Video
	var authorList []uint
	numVideo, err := service.GetNumVideo(&videoList, &authorList, LastTime, global.MaxNumVideo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": -1, "status_msg": "Fail to get the number of video."})
		return
	}
	// 如果是空的
	if numVideo == 0 {
		// 没有满足条件的视频
		c.JSON(http.StatusOK, feedResponse{
			Response:  db.Response{StatusCode: 0, StatusMsg: "null"},
			VideoList: nil,
			NextTime:  CurrentTime,
		})
		return
	}

	// 视频加入视频列表内，互动接口功能后续实现
	// 初始化信息，后续完善了互动功能再补充
	var (
		videoJsonList []db.Video
		videoJson     db.Video
	)

	// 省略社交功能
	for _, video := range videoList {

		videoJson.Videoid = video.Videoid
		videoJson.Title = video.Title
		videoJson.Userid = video.Userid
		videoJson.Videourl = "http://" + c.Request.Host + "/static/video/" + video.Videourl
		videoJson.Coverurl = "http://" + c.Request.Host + "/static/cover/" + video.Coverurl

		videoJsonList = append(videoJsonList, videoJson)
	}

	nextTime := videoList[numVideo-1].CreatedAt.Unix()
	// 输出视频流
	c.JSON(http.StatusOK, feedResponse{
		Response:  db.Response{StatusCode: 0, StatusMsg: "null"},
		VideoList: videoJsonList,
		NextTime:  nextTime,
	})
}
