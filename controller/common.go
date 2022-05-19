package controller

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primary_key"`
	Author        User   `json:"author" gorm:"foreignKey:id;references:id;"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}
type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key"`
	User       User   `json:"user" gorm:"foreignKey:id;references:id;"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type UserVideoFavorite struct {
	VideoId    int64
	UserId     int64
	IsFavorite bool
}
type UserUserFollow struct {
	Id       int64
	FansId   int64
	IsFollow bool
}
type UserLogin struct {
	Id       int64
	Name     string
	Password string
}
