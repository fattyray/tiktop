package service

import (
	"fmt"
	"os/exec"
	"tiktop/entity"
	"tiktop/global"
)

// GetCoverFromVideo 根据视频生成封面图片
func GetCoverFromVideo(videoPath, coverPath string) error {
	cmd := exec.Command("ffmpeg",
		"-i", videoPath, "-r", "1",
		"-vframes", "1",
		"-f", "image2",
		coverPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("生成封面失败：%s\n%s", err, output)
	}

	return nil
}

// 查询视频id列表
func QueryVideoIdListByUserId(userId int64) (videoIdList []int64, err error) {
	result := global.DB.Model(&entity.Video{}).Select("video_id").Where("user_id = ?", userId).Find(&videoIdList)
	if result.Error != nil {
		err = result.Error
		return nil, err
	}
	return
}

// 查询视频对象列表
func QueryVideoListByUserId(userId int64) (videoList []entity.Video, err error) {
	if global.DB.Where("user_id = ?", userId).Find(&videoList).Error != nil {
		return
	}
	return
}

// 查询视频封装返回对象列表
func GetPostVideoListByUserId(userId int64) (videos []entity.VideoResponse, err error) {
	//查询视频对象列表
	videoList, err := QueryVideoListByUserId(userId)
	if err != nil {
		return nil, err
	}
	//构造视频id列表
	videoIdList := make([]int64, len(videoList))
	for i, video := range videoList {
		videoIdList[i] = video.VideoId
	}
	//根据视频id列表查询点赞数量
	likeCountList, err := QueryLikeCountListByVideoIdList(&videoIdList)
	if err != nil {
		return nil, err
	}
	likeCountListMap := map[int64]int64{}
	for _, likeCount := range likeCountList {
		likeCountListMap[likeCount.VideoId] = likeCount.LikeCnt
	}
	//根据视频id列表查询评论数量
	commentCountList, err := QueryCommentCountListByVideoIdList(&videoIdList)
	if err != nil {
		return nil, err
	}
	commentCountListMap := map[int64]int64{}
	for _, likeCount := range commentCountList {
		commentCountListMap[likeCount.VideoId] = likeCount.CommentCnt
	}
	videos = make([]entity.VideoResponse, len(videoList))
	for i, video := range videoList {
		videos[i].Id = video.VideoId
		videos[i].Author, err = UserInfoByUserId(video.UserId)
		if err != nil {
			return nil, err
		}
		//仅有登录用户自己
		videos[i].Author.IsFollow, err = QueryFollowOrNot(userId, userId)
		if err != nil {
			return nil, err
		}
		videos[i].CoverUrl = video.CoverUrl
		videos[i].PlayUrl = video.PlayUrl
		videos[i].Title = video.Title
		videos[i].IsFavorite = true
		videos[i].FavoriteCount = likeCountListMap[video.VideoId]
		videos[i].CommentCount = commentCountListMap[video.VideoId]
	}
	return
}
