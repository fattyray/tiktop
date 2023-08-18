package service

import "tiktop/entity"

// 在一个给定的 int64 数组中查找给定元素
func FindInt64(target int64, intArr *[]int64) bool {
	for _, element := range *intArr {
		if target == element {
			return true
		}
	}
	return false
}

// 在一个给定 VideoLikeCnt 列表中查找给定视频 id 是否存在，不存在返回 0，存在返回点赞值
func FindVideoIdFromVideoLikeCntList(videoId int64, likeCountList *[]entity.VideoLikeCnt) int64 {
	for _, element := range *likeCountList {
		if videoId == element.VideoId {
			return element.LikeCnt
		}
	}
	return 0
}

// 在一个给定 VideoCommentCnt 列表中查找给定视频 id 是否存在，不存在返回 0，存在返回评论值
func FindVideoIdFromVideoCommentCntList(videoId int64, commentCountList *[]entity.VideoCommentCnt) int64 {
	for _, element := range *commentCountList {
		if videoId == element.VideoId {
			return element.CommentCnt
		}
	}
	return 0
}
