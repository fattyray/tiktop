package service

import (
	"errors"
	"tiktop/entity"
	"tiktop/global"
	"tiktop/util"
	"time"
)

// 根据视频id列表查询评论数量列表
func QueryCommentCountListByVideoIdList(videoIdList *[]int64) ([]entity.VideoCommentCnt, error) {
	var getCommentCountList []entity.VideoCommentCnt
	result := global.DB.Model(&entity.Comment{}).Select("video_id", "count(video_id) as comment_cnt").Where("video_id in ?", *videoIdList).Group("video_id").Find(&getCommentCountList)
	if result.Error != nil {
		err := errors.New("commentList query failed")
		return nil, err
	}
	// 找数据找齐了
	if len(*videoIdList) == len(getCommentCountList) {
		return getCommentCountList, nil
	}
	// 数据不全，误差部分补全为0
	var commentCountList []entity.VideoCommentCnt
	commentCountList = make([]entity.VideoCommentCnt, len(*videoIdList))
	for i, videoId := range *videoIdList {
		commentCountList[i].VideoId = videoId
		commentCountList[i].CommentCnt = FindVideoIdFromVideoCommentCntList(videoId, &getCommentCountList)
	}
	return commentCountList, nil
}

// 增加评论
func AddComment(currentId int64, videoId int64, commentText string) (err error) {
	var addComment entity.Comment
	addComment.CommentId = util.GetNextId()
	addComment.UserId = currentId
	addComment.VideoId = videoId
	addComment.Content = commentText
	addComment.CreateDate = time.Now().Format("01-02")
	result := global.DB.Model(&entity.Comment{}).Create(&addComment)
	if result.Error != nil {
		return err
	}
	return
}

// 删除评论
func CancelComment(currentId int64, videoId int64, commentId int64) (err error) {
	var cancelComment entity.Comment
	cancelComment.CommentId = commentId
	cancelComment.UserId = currentId
	cancelComment.VideoId = videoId
	result := global.DB.Model(&entity.Comment{}).Delete(&cancelComment)
	if result.Error != nil {
		return err
	}
	if result.RowsAffected == 0 {
		err = errors.New("comment not found")
		return err
	}
	return
}

func GetCommentListByVideoId(currentId int64, videoId int64) (comments []entity.CommentResponse, err error) {
	var commentList []entity.Comment
	if global.DB.Model(&entity.Comment{}).Where("video_id=?", videoId).Find(&commentList).Error != nil {
		return
	}
	comments = make([]entity.CommentResponse, len(commentList))
	for i, comment := range commentList {
		comments[i].Id = comment.CommentId
		comments[i].User, err = UserInfoByUserId(comment.UserId)
		if err != nil {
			return nil, err
		}
		comments[i].User.IsFollow, err = QueryFollowOrNot(currentId, comment.UserId)
		if err != nil {
			return nil, err
		}
		comments[i].Content = comment.Content
		comments[i].CreateDate = comment.CreateDate
	}
	return
}
