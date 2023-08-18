package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktop/entity"
	"tiktop/global"
	"tiktop/service"
	"tiktop/util"
	"time"
)

type feedResponse struct {
	Response  entity.Response
	VideoList []entity.VideoResponse `json:"video_list,omitempty"`
	NextTime  int64                  `json:"next_time,omitempty"`
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
	var videoList []entity.Video
	var videoIdList []int64
	var authorIdList []int64
	numVideo, err := service.GetNumVideo(&videoList, &videoIdList, &authorIdList, LastTime, global.MaxNumVideo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": -1, "status_msg": "Fail to get the number of video."})
		return
	}
	//fmt.Println(numVideo)

	// 如果是空的
	if numVideo == 0 {
		// 没有满足条件的视频
		c.JSON(http.StatusOK, feedResponse{
			Response:  entity.Response{StatusCode: 0, StatusMsg: "null"},
			VideoList: nil,
			NextTime:  CurrentTime,
		})
		return
	}

	// 点赞信息获得
	LikeVideoList, errLike := service.QueryLikeCountListByVideoIdList(&videoIdList)
	if errLike != nil {
		c.JSON(http.StatusNotFound, feedResponse{
			Response:  entity.Response{StatusCode: 1, StatusMsg: "Fail to get liked count for videos."},
			VideoList: nil,
			NextTime:  LastTime,
		})
		return
	}
	//fmt.Println(LikeVideoList)

	// 评论信息获得
	CommentVideoList, errComment := service.QueryCommentCountListByVideoIdList(&videoIdList)
	if errComment != nil {
		c.JSON(http.StatusNotFound, feedResponse{
			Response:  entity.Response{StatusCode: 1, StatusMsg: "Fail to get comment count for videos."},
			VideoList: nil,
			NextTime:  LastTime,
		})
		return
	}
	//fmt.Println(CommentVideoList)

	// 点赞与否
	// 登录状态判断
	var userid int64
	isLogged := false
	token := c.PostForm("token")
	if token == "" {
		token = c.Query("token")
	}
	if token != "" {
		claims, errToken := util.Gettoken(token)
		if errToken == nil {
			userid, err = strconv.ParseInt(claims.UserId, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, feedResponse{
					Response:  entity.Response{StatusCode: 1, StatusMsg: err.Error()},
					VideoList: nil,
					NextTime:  LastTime,
				})
				return
			}

			isLogged = true
		}
	}

	// 点赞与关注判断
	var isFavoriteList []bool
	var isFollowList []bool

	if isLogged {
		// 点赞列表
		isFavoriteList, err = service.ParseLikeVideoListByUserIdFormVideoId(userid, &videoIdList)
		if err != nil {
			c.JSON(http.StatusBadRequest, feedResponse{
				Response:  entity.Response{StatusCode: 1, StatusMsg: err.Error()},
				VideoList: nil,
				NextTime:  LastTime,
			})
			return
		}
		// 关注列表
		isFollowList, err = service.ParseFollowListByUserIdFormUserId(userid, &authorIdList)
		if err != nil {
			c.JSON(http.StatusBadRequest, feedResponse{
				Response:  entity.Response{StatusCode: 1, StatusMsg: err.Error()},
				VideoList: nil,
				NextTime:  LastTime,
			})
			return
		}
	}

	isFavorite := false

	// 初始化列表信息
	var (
		videoJsonList []entity.VideoResponse
		videoJson     entity.VideoResponse
		author        entity.UserData
	)

	// 填充输出信息
	for i, video := range videoList {
		// author 获取
		author, err = service.UserInfoByUserId(authorIdList[i])
		if err != nil {
			//fmt.Println("Not found user")
			c.JSON(http.StatusNotFound, feedResponse{
				Response:  entity.Response{StatusCode: 1, StatusMsg: err.Error()},
				VideoList: nil,
				NextTime:  LastTime,
			})
			return
		}

		// 登录时信息获取
		if isLogged {
			author.IsFollow = isFollowList[i]
			isFavorite = isFavoriteList[i]
		}

		// 信息填充
		videoJson.Id = video.VideoId
		videoJson.Author = author
		videoJson.PlayUrl = video.PlayUrl
		videoJson.CoverUrl = video.CoverUrl
		videoJson.FavoriteCount = LikeVideoList[i].LikeCnt
		videoJson.CommentCount = CommentVideoList[i].CommentCnt
		videoJson.IsFavorite = isFavorite
		videoJson.Title = video.Title

		videoJsonList = append(videoJsonList, videoJson)
	}

	nextTime := videoList[numVideo-1].CreatedAt.Unix()
	// 输出视频流
	c.JSON(http.StatusOK, feedResponse{
		Response:  entity.Response{StatusCode: 0, StatusMsg: "null"},
		VideoList: videoJsonList,
		NextTime:  nextTime,
	})
}
