package service

import (
	"log"
	"tiktop/entity"
	"tiktop/global"
)

func AddMessage(fromUserId int64, toUserId int64, content string) (info string, err error) {
	message := entity.Message{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		Content:    content,
	}
	//插入数据库
	if err := global.DB.Create(&message).Error; err != nil {
		log.Print("向message数据库中插入数据失败！")
		log.Println(err)
		return "向message数据库中插入数据失败！", err
	}
	//没出错
	return "向message数据库中插入成功", nil

}
func ChatList(fromUserId int64, toUserId int64, msgTime int64) (messageList []entity.Message, err error) {
	if msgTime == 0 {
		if err := global.DB.Where("  ((`from_user_id` = ? AND `to_user_id` = ?)  OR (`from_user_id` = ? AND `to_user_id` = ?))", fromUserId, toUserId, toUserId, fromUserId).Find(&messageList).Error; err != nil {
			log.Print("从message数据库中查询数据失败！")
			log.Println(err)
			return messageList, err
		}
	} else {
		if err := global.DB.Where("create_time > ? AND `from_user_id` = ? AND `to_user_id` = ?", msgTime, toUserId, fromUserId).Find(&messageList).Error; err != nil {
			log.Print("从message数据库中查询数据失败！")
			log.Println(err)
			return messageList, err
		}
	}
	//查询数据库
	if err := global.DB.Where("create_time > ? AND ((`from_user_id` = ? AND `to_user_id` = ?)  OR (`from_user_id` = ? AND `to_user_id` = ?))", msgTime, fromUserId, toUserId, toUserId, fromUserId).Find(&messageList).Error; err != nil {
		log.Print("从message数据库中查询数据失败！")
		log.Println(err)
		return messageList, err
	}
	conunt := 0
	//没出错
	//如果前N条是自己发的消息并且不是首次获取消息就把自己发的消息删除
	for i := range messageList {
		if messageList[i].FromUserId != fromUserId || msgTime == 0 {
			break
		} else {
			conunt++
		}
	}
	if conunt != 0 {
		messageList = messageList[conunt:]
	}

	return messageList, nil
}
