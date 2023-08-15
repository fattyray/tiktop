package service

import (
	"errors"
	"tiktop/entity"
	"tiktop/global"
	"tiktop/util"
)

// 获取每个视频的点赞数量
func QueryLikeCountListByVideoIdList(videoIdList *[]int64) (likeCountList []entity.VideoLikeCnt, err error) {
	result := global.DB.Model(&entity.Like{}).Select("video_id", "count(video_id) as like_cnt").Where(map[string]interface{}{"video_id": *videoIdList}).Group("video_id").Find(&likeCountList)
	if result.Error != nil {
		err = errors.New("likesList query failed")
		return
	}
	return
}

// 使用用户id查询其点赞视频的id列表
func QueryLikeVideoIdListByUserId(userId int64) (likeList []int64, err error) {
	result := global.DB.Model(&entity.Like{}).Select("video_id").Where("user_id=?", userId).Find(&likeList)
	if result.Error != nil {
		return nil, err
	}
	return
}

// 根据视频id查询视频对象
func QueryVideoListByVideoIdList(videoIdList *[]int64) (videoList []entity.Video, err error) {
	result := global.DB.Model(&entity.Video{}).Where("video_id in ?", *videoIdList).Find(&videoList)
	if result.Error != nil {
		return nil, err
	}
	return
}

// 点赞操作和取消赞操作
func GiveOrCancelLike(userId int64, videoId int64, actionType int32) (err error) {
	var likeList []entity.Like
	result := global.DB.Model(&entity.Like{}).Where("user_id=? and video_id=?", userId, videoId).Find(&likeList)
	if result.Error != nil {
		return
	}
	//查询到有点赞记录
	if likeList != nil && len(likeList) > 0 {
		//已经点赞过
		if actionType == 1 {
			return
		}
		//取消点赞
		var cancelLike entity.Like
		cancelLike.LikeId = likeList[0].LikeId
		result = global.DB.Model(&entity.Like{}).Delete(&cancelLike)
		if result.Error != nil {
			return err
		}
		return
	}
	//无点赞记录
	//取消点赞
	if actionType == 2 {
		return
	}
	//进行点赞
	var giveLike entity.Like
	giveLike.LikeId = util.GetNextId()
	giveLike.UserId = userId
	giveLike.VideoId = videoId
	if global.DB.Model(&entity.Like{}).Create(&giveLike).Error != nil {
		return err
	}
	return
}

// 根据id查询点赞视频列表
func GetLikeVideoListByUserId(userId int64, currentId int64) (videos []entity.VideoResponse, err error) {
	//查询当前用户的点赞的视频id列表
	likeVideoIdList, err := QueryLikeVideoIdListByUserId(currentId)
	if err != nil {
		return nil, err
	}
	//根据视频id列表查询视频对象
	likeVideoList, err := QueryVideoListByVideoIdList(&likeVideoIdList)
	if err != nil {
		return nil, err
	}
	//根据视频id列表查询点赞数量
	likeCountList, err := QueryLikeCountListByVideoIdList(&likeVideoIdList)
	if err != nil {
		return nil, err
	}
	//防止数量为0,预先使用map记录
	likeCountListMap := map[int64]int64{}
	for _, likeCount := range likeCountList {
		likeCountListMap[likeCount.VideoId] = likeCount.LikeCnt
	}
	//根据视频id列表查询评论数量
	commentCountList, err := QueryCommentCountListByVideoIdList(&likeVideoIdList)
	if err != nil {
		return nil, err
	}
	commentCountListMap := map[int64]int64{}
	for _, likeCount := range commentCountList {
		commentCountListMap[likeCount.VideoId] = likeCount.CommentCnt
	}
	videos = make([]entity.VideoResponse, len(likeVideoList))
	for i, video := range likeVideoList {
		videos[i].Id = video.VideoId
		videos[i].Author, err = UserInfoByUserId(video.UserId)
		if err != nil {
			return nil, err
		}
		videos[i].Author.IsFollow, err = QueryFollowOrNot(currentId, userId)
		if err != nil {
			return nil, err
		}
		videos[i].CoverUrl = video.CoverUrl
		videos[i].PlayUrl = video.PlayUrl
		videos[i].Title = video.Title
		videos[i].IsFavorite = true
		//map中没有数据则自动为0
		videos[i].FavoriteCount = likeCountListMap[video.VideoId]
		videos[i].CommentCount = commentCountListMap[video.VideoId]
	}

	return
}
