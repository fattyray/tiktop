package service

import (
	"errors"
	"github.com/bwmarrin/snowflake"
	"tiktop/entity"
	"tiktop/global"
)

func IsUserIdExist(userId int64) (flag bool, err error) {
	var user []entity.User
	result := global.DB.Where("user_id = ?", userId).Find(&user)
	if user == nil || len(user) == 0 || result.Error != nil {
		err = errors.New("user not found")
		return false, err
	}
	return true, nil
}

func Register(username string, password string) (user *entity.User, err error) {
	//判断用户名是否存在
	result := global.DB.Model(&entity.User{}).Where("user_name = ?", username).First(&user)
	if result.RowsAffected != 0 {
		err = errors.New("username already exists")
		return
	}
	user.UserName = username
	user.Password = password
	node, err := snowflake.NewNode(1)
	if err != nil {
		err = errors.New("userID generate failed")
		return
	}
	user.UserId = node.Generate().Int64()
	err = global.DB.Create(user).Error //存储到数据库
	return
}
func Login(username string, password string) (user *entity.User, err error) {
	//检查用户名是否存在
	result := global.DB.Model(&entity.User{}).Where("user_name = ? and password= ? ", username, password).First(&user)
	if result.RowsAffected == 0 {
		err = errors.New("login failed")
		return
	}
	return
}
func UserInfoByUserId(userId int64) (userdata entity.UserData, err error) {
	var user entity.User
	result := global.DB.Where("user_id = ?", userId).First(&user)
	if result.RowsAffected == 0 {
		err = errors.New("user not found")
		return
	}
	userdata.UserId = userId
	userdata.Name = user.UserName
	userdata.Avatar = user.Avatar
	userdata.BackgroundImage = user.BackgroundImage
	userdata.Signature = user.Signature
	global.DB.Model(&entity.Follow{}).Where("user_id = ?", userId).Count(&userdata.FollowCount)
	global.DB.Model(&entity.Follow{}).Where("follow_userid = ?", userId).Count(&userdata.FollowerCount)
	global.DB.Model(&entity.Like{}).Where("user_id = ?", userId).Count(&userdata.FavoriteCount)
	videoIdList, err1 := QueryVideoIdListByUserId(userId)
	if err1 != nil {
		err = err1
		return
	}
	userdata.WorkCount = int64(len(videoIdList))
	likesList, err2 := QueryLikeCountListByVideoIdList(&videoIdList)
	if err2 != nil {
		err = err2
		return
	}
	var totalCnt int64 = 0
	for _, like := range likesList {
		totalCnt += like.LikeCnt
	}
	userdata.TotalFavorited = totalCnt

	// 默认未登录为 false, 交由接口处理
	userdata.IsFollow = false

	return
}
