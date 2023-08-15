package views

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"strconv"
	"tiktop/global"
)

package views

import (
"fmt"
"github.com/bwmarrin/snowflake"
"github.com/gin-gonic/gin"
"gorm.io/driver/mysql"
"gorm.io/gorm"
"net/http"
"path/filepath"
"strconv"
"tiktop/db"
"tiktop/global"
"tiktop/service"
"tiktop/sql_dsn"
)

func PublishAction(c *gin.Context) {
	// 获取视频信息
	userid := c.GetUint("user_id")
	fmt.Println(userid)
	title := c.PostForm("title")
	fmt.Println(title)

	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusInternalServerError, db.Response{
			StatusCode: -1,
			StatusMsg:  "fail to upload the file.",
		})
		return
	}

	// 读取数据库
	dsn := sql_dsn.GetDsn()
	dbx, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusBadRequest, db.Response{
			StatusCode: -1,
			StatusMsg:  "failed to connect database",
		})
		return
	}

	// 获取视频唯一标识 id
	node, err := snowflake.NewNode(1)
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "failed to generate snowflake"})
		c.JSON(http.StatusBadRequest, db.Response{
			StatusCode: 1,
			StatusMsg:  "failed to generate snowflake for video",
		})
	}
	videoId := uint(node.Generate().Int64())

	// 获取视频路径，并且把视频以及封面存放入指定位置
	name := strconv.FormatUint(uint64(videoId), 10)
	//filename :=
	videoName := name + file.Filename
	coverName := name + ".jpg"

	videoSavePath := filepath.Join(global.PATH_VIDEO, videoName)
	coverSavePath := filepath.Join(global.PATH_COVER, coverName)

	err = c.SaveUploadedFile(file, videoSavePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, db.Response{
			StatusCode: -1,
			StatusMsg:  "fail to save the file to the path.",
		})
		return
	}

	err = service.GetCoverFromVideo(videoSavePath, coverSavePath)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, db.Response{
			StatusCode: -1,
			StatusMsg:  "fail to create the cover.",
		})
		return
	}

	// 把视频信息生成 video 结构体，并且存入数据库
	video := db.Video{
		Videoid:  videoId,
		Videourl: videoSavePath,
		Coverurl: coverSavePath,
		Title:    title,
		Userid:   userid,
	}
	err = dbx.Create(&video).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, db.Response{
			StatusCode: -1,
			StatusMsg:  "fail to add the video into SQL",
		})
		return
	}

	fmt.Printf("视频写入数据库\n")

	c.JSON(http.StatusOK, db.Response{
		StatusCode: 0,
		StatusMsg:  "null",
	})
}