package entity

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type FavoriteListResponse struct {
	Response
	VideoList []VideoResponse `json:"video_list"`
}
type VideoResponse struct {
	Id            int64    `json:"id"`
	Author        UserData `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

//type UserInf struct {
//	Id            int64  `json:"id,omitempty"`
//	Name          string `json:"name,omitempty"`
//	FollowCount   int64  `json:"follow_count,omitempty"`
//	FollowerCount int64  `json:"follower_count,omitempty"`
//	IsFollow      bool   `json:"is_follow,omitempty"`
//}

type CommentActionResponse struct {
	Response        Response
	CommentResponse CommentResponse
}
type CommentResponse struct {
	Id         int64    `json:"id"`
	User       UserData `json:"user"`
	Content    string   `json:"content"`
	CreateDate string   `json:"create_date"`
}
type CommentListResponse struct {
	Response Response
	Comments []CommentResponse `json:"comment_list"`
}

type FollowListResponse struct {
	Response Response
	UserList []UserData `json:"user_list"`
}
