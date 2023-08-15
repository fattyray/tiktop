package entity

type Follow struct {
	FollowId     int64 `gorm:"column:follow_id;primary_key;NOT NULL"`
	UserId       int64 `gorm:"column:user_id;NOT NULL"`
	FollowUserId int64 `gorm:"column:follow_userid;NOT NULL"`
}
