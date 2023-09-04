package entity

import (
	"fmt"
	"time"
)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}
type Message struct {
	Id         int64  `json:"id"  gorm:"primaryKey"`
	ToUserId   int64  `json:"to_user_id" `
	FromUserId int64  `json:"from_user_id" `
	Content    string `json:"content" gorm:"type:varchar(200)"`
	CreateTime int64  `json:"create_time" gorm:"autoCreateTime:milli;index:;sort:desc"`
}
type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}
