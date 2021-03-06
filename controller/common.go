package controller

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primary_key"`
	UserId        int64  `json:"user_id,omitempty" gorm:"primary_key"`
	Author        User   `json:"author,omitempty" gorm:"foreignKey:UserId;references:Id;"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64 `json:"id,omitempty" gorm:"primary_key"`
	VideoId    int64
	UserId     int64  `json:"user_id,omitempty" `
	User       User   `json:"user" gorm:"foreignKey:UserId;references:id;"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
}

type UserVideoFavorite struct {
	VideoId    int64 `gorm:"primary_key"`
	UserId     int64 `gorm:"primary_key"`
	IsFavorite bool  `json:"is_favorite,omitempty"`
}

type UserUserFollow struct {
	Id     int64 `gorm:"primary_key"`
	FansId int64 `gorm:"primary_key"`
}

type UserLogin struct {
	Id       int64 `gorm:"primary_key"`
	Name     string
	Password string
}
