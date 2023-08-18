package service

import (
	"tiktop/entity"
	"tiktop/global"
)

// GetNumVideo 获取视频列表中符合结果的视频数据，以及相应的视频，作者信息列表
func GetNumVideo(videos *[]entity.Video, videoIdList *[]int64, AuthorIdList *[]int64, LastTime int64, MaxNumVide int) (int, error) {
	query := global.DB.Order("created_at desc").
		Limit(MaxNumVide).
		Where("UNIX_TIMESTAMP(created_at) <= ?", LastTime)
	query.Find(videos)

	numVideo := len(*videos)

	// 统计作者 id 以及视频 id
	*AuthorIdList = make([]int64, numVideo)
	*videoIdList = make([]int64, numVideo)
	for i, video := range *videos {
		(*AuthorIdList)[i] = video.UserId
		(*videoIdList)[i] = video.VideoId
	}

	return numVideo, nil
}
