package entity

type Like struct {
	LikeId  int64 `gorm:"column:like_id;primary_key;NOT NULL"`
	UserId  int64 `gorm:"column:user_id;NOT NULL"`
	VideoId int64 `gorm:"column:video_id;NOT NULL"`
}
