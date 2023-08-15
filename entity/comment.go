package entity

type Comment struct {
	CommentId  int64  `gorm:"column:comment_id;primary_key;NOT NULL"`
	UserId     int64  `gorm:"column:user_id;NOT NULL"`
	VideoId    int64  `gorm:"column:video_id;NOT NULL"`
	Content    string `gorm:"column:comment_text;NOT NULL;type:varchar(300)"`
	CreateDate string `gorm:"column:create_date;NOT NULL;type:varchar(100)"`
}
