package service

import (
	"errors"
	"sort"
	"tiktop/entity"
	"tiktop/global"
	"tiktop/util"
)

// 查询是否关注
func QueryFollowOrNot(userId int64, toId int64) (res bool, err error) {
	var follow []entity.Follow
	result := global.DB.Model(&entity.Follow{}).Where("user_id = ? and follow_userid = ?", userId, toId).Find(&follow)
	if result.RowsAffected == 0 {
		return false, nil
	}
	if result.Error != nil {
		err = result.Error
		return
	}
	return true, nil
}

// 关注和取消关注
func DoOrCancelFollow(currentId int64, userId int64, actionType int32) (err error) {
	//判断目标用户id是否存在
	flag, err := IsUserIdExist(userId)
	if err != nil || !flag {
		err = errors.New("user not found")
		return
	}

	//查询是否已经关注
	var follow []entity.Follow
	if global.DB.Model(&entity.Follow{}).Where("user_id=? and follow_userid=?", currentId, userId).Find(&follow).Error != nil {
		err = errors.New("follow query failed")
		return err
	}
	if actionType == 1 {
		//重复关注
		if follow != nil && len(follow) > 0 {
			return nil
		}
		//关注
		var followact entity.Follow
		followact.FollowId = util.GetNextId()
		followact.UserId = currentId
		followact.FollowUserId = userId
		if global.DB.Model(&entity.Follow{}).Create(&followact).Error != nil {
			err = errors.New("do follow failed")
			return err
		}
	} else if actionType == 2 {
		if follow == nil || len(follow) == 0 {
			return nil
		}
		if global.DB.Model(&entity.Follow{}).Delete(&follow[0]).Error != nil {
			err = errors.New("cancel follow failed")
			return err
		}
	}
	return
}

// 查询关注用户id列表
func getFollowUserIdListByUserId(userId int64) (followUserIdList []int64, err error) {
	if global.DB.Model(&entity.Follow{}).Select("follow_userid").Where("user_id=?", userId).Find(&followUserIdList).Error != nil {
		err = errors.New("follow userid query failed")
		return nil, err
	}
	return
}

// 根据用户id以及给定作者id列表返回关注列表情况
func ParseFollowListByUserIdFormUserId(userId int64, authorIdList *[]int64) (isFollowList []bool, err error) {
	var followUserIdList []int64
	followUserIdList, err = getFollowUserIdListByUserId(userId)
	if err != nil {
		return nil, err
	}
	sort.Slice(followUserIdList, func(i, j int) bool { return followUserIdList[i] < followUserIdList[j] })
	isFollowList = make([]bool, len(*authorIdList))
	for i, authorId := range *authorIdList {
		isFollowList[i] = FindInt64(authorId, &followUserIdList)
	}
	return
}

// 查询关注用户列表
func GetFollowUserListByUserId(userId int64) (followUserList []entity.UserData, err error) {
	followUserIdList, err := getFollowUserIdListByUserId(userId)
	if err != nil {
		return nil, err
	}
	followUserList = make([]entity.UserData, len(followUserIdList))
	for i, followUserId := range followUserIdList {
		followUserList[i], err = UserInfoByUserId(followUserId)
		if err != nil {
			return nil, err
		}
		followUserList[i].IsFollow = true
	}
	return
}

// 查询粉丝id列表
func getFollowerIdListByUserId(userId int64) (followerIdList []int64, err error) {
	if global.DB.Model(&entity.Follow{}).Select("user_id").Where("follow_userid=?", userId).Find(&followerIdList).Error != nil {
		err = errors.New("follow userid query failed")
		return nil, err
	}
	return
}

// 查询粉丝列表
func GetFollowerListByUserId(userId int64) (followerList []entity.UserData, err error) {
	followerIdList, err := getFollowerIdListByUserId(userId)
	if err != nil {
		return nil, err
	}
	followerList = make([]entity.UserData, len(followerIdList))
	for i, followerId := range followerIdList {
		followerList[i], err = UserInfoByUserId(followerId)
		if err != nil {
			return nil, err
		}
		followerList[i].IsFollow, err = QueryFollowOrNot(userId, followerId)
		if err != nil {
			return nil, err
		}
	}
	return
}

//func GetFriendListByUserId(userId int64) (userList []entity.UserData, err error) {
//	followerList, err := GetFollowerListByUserId(userId)
//	if err != nil {
//		return nil, err
//	}
//	for _, follower := range followerList {
//		flag, err := QueryFollowOrNot(userId, follower.UserId)
//		if err != nil {
//			return nil, err
//		}
//		if flag {
//			userInfo, err := UserInfoByUserId(follower.UserId)
//			userInfo.IsFollow = true
//			if err != nil {
//				return nil, err
//			}
//			userList = append(userList, userInfo)
//		}
//	}
//	return
//}
