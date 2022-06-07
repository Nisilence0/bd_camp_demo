package controller

import (
	"fmt"

	"gorm.io/gorm/clause"
)

func GetAllVideos() (bool, error) {

	err := db.Preload(clause.Associations).Find(&DemoVideos).Error

	if err != nil {
		return false, err
	}
	fmt.Printf("%v", DemoVideos)
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
		CoverUrl:      "http://localhost/public/cover.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}).Error

	if err != nil {
		return false, err
	}

	_, err = GetAllVideos()
	if err != nil {
		return false, err
	}

	return true, nil
}

func SearchUserVideo(user User) (bool, error) {

	err := db.Where("user_id = ?", user.Id).Find(&PublUserVideo).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
