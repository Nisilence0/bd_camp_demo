package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetAllVideos(c *gin.Context) (bool, error) {
	
	db.Model(&Video{}).Where("is_favorite = ?", 1).Update("is_favorite", 0)

	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {

		db.Model(&Video{}).Update("is_favorite", 0)

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

func GetAllVideos2() (bool, error) {

	err := db.Preload(clause.Associations).Find(&DemoVideos).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func GetVideosCount() (int64, error) {

	var count int64

	err := db.Preload(clause.Associations).Find(&DemoVideos).Count(&count).Error

	if err != nil {
		return -1, err
	}
	return count, nil

}

func AddVideos(user User, playurl string) (bool, error) {

	err := db.Create(&Video{
		UserId:        user.Id,
		PlayUrl:       playurl,
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

func SearchUserVideo(user User) (bool, error) {

	err3 := db.Where("user_id = ?", user.Id).Find(&PublUserVideo).Error

	if err3 != nil {
		return false, err3
	}

	return true, nil
}
