package entity

type UserData struct {
	UserId          int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type User struct {
	UserId          int64  `gorm:"column:user_id;primary_key;NOT NULL"`
	UserName        string `gorm:"column:user_name;type:varchar(100)"`
	Password        string `gorm:"column:password;type:varchar(100)"`
	Avatar          string `gorm:"column:avatar;type:varchar(100)"`
	BackgroundImage string `gorm:"column:background_image;type:varchar(100)"`
	Signature       string `gorm:"column:signature;type:varchar(200)"`
}
