package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

//获取以用户点赞为基准的视频列表
func GetAllVideos(c *gin.Context) (bool, error) {

	ClearVideoFavorite()

	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {

		err := db.Model(&Video{}).Where("id in (?)", db.Where("user_id = ?", user.Id).Select("video_id").Find(&FavoUserVideoList)).Update("is_favorite", true).Error

		if err != nil {
			return false, err
		}

	}

	err := db.Preload(clause.Associations).Find(&DemoVideos).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

//获取视频列表
func GetAllVideos2() (bool, error) {

	err := db.Preload(clause.Associations).Find(&DemoVideos).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

//用户上传视频
func AddVideos(user User, playurl string) (bool, error) {

	err := db.Create(&Video{
		UserId:        user.Id,
		PlayUrl:       "http://150.158.197.247:80/public/" + playurl,
		CoverUrl:      "http://150.158.197.247:80/public/cover.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}).Error

	if err != nil {
		return false, err
	}

	_, err = GetAllVideos2()
	if err != nil {
		return false, err
	}

	return true, nil
}

//查找用户发表的视频
func SearchUserVideo(user User) (bool, error) {

	err3 := db.Where("user_id = ?", user.Id).Find(&PublUserVideo).Error

	if err3 != nil {
		return false, err3
	}

	return true, nil
}
