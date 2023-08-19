package controller

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"tiktop/entity"
	"tiktop/global"
	"tiktop/service"
	"tiktop/util"
)

type videoListResponse struct {
	entity.Response
	VideoList []entity.VideoResponse `json:"video_list"`
}

func Publish(c *gin.Context) {
	//校验token并获取当前用户id
	token := c.PostForm("token")
	claims, _ := util.Gettoken(token)
	userid, _ := strconv.ParseInt(claims.UserId, 10, 64)
	title := c.PostForm("title")

	// 获取文件
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: -1,
			StatusMsg:  "fail to upload the file.",
		})
		return
	}

	// 获取视频唯一标识 id
	node, err := snowflake.NewNode(1)
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "failed to generate snowflake"})
		c.JSON(http.StatusBadRequest, entity.Response{
			StatusCode: 1,
			StatusMsg:  "failed to generate snowflake for video",
		})
	}
	videoId := node.Generate().Int64()

	// 获取视频路径，并且把视频以及封面存放入指定位置
	name := strconv.FormatUint(uint64(videoId), 10)
	//filename :=
	videoName := name + file.Filename
	coverName := name + ".jpg"

	videoSavePath := filepath.Join(global.PATH_VIDEO, videoName)
	//coverSavePath := filepath.Join(global.PATH_COVER, coverName)

	err = c.SaveUploadedFile(file, videoSavePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{
			StatusCode: -1,
			StatusMsg:  "fail to save the file to the path.",
		})
		return
	}

	//err = service.GetCoverFromVideo(videoSavePath, coverSavePath)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, entity.Response{
			StatusCode: -1,
			StatusMsg:  "fail to create the cover.",
		})
		return
	}

	// 把视频信息生成 video 结构体，并且存入数据库
	playUrl := global.HEAD_URL + c.Request.Host + global.VIDEO_URL + videoName
	coverUrl := global.HEAD_URL + c.Request.Host + global.COVER_URL + coverName
	video := entity.Video{
		VideoId:  videoId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
		UserId:   userid,
	}
	err = global.DB.Create(&video).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: -1,
			StatusMsg:  "fail to add the video into SQL",
		})
		return
	}

	//fmt.Printf("视频写入数据库\n")

	c.JSON(http.StatusOK, entity.Response{
		StatusCode: 0,
		StatusMsg:  "null",
	})
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
