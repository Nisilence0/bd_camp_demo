package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//清除当前视频已点赞
func ClearVideoFavorite() (bool, error) {

	if err := db.Model(&Video{}).Where("is_favorite = ?", 1).Update("is_favorite", 0).Error; err != nil {
		return false, err
	}

	return true, nil

}

// FavoriteAction no practical effect, just check if token is valid
//储存用户点赞视频
func SaveUserFavoriteVideo(user User, video_id int64) (bool, error) {

	if err := db.Create(&UserVideoFavorite{UserId: user.Id, VideoId: video_id, IsFavorite: true}).Error; err != nil {
		return false, err
	}
	db.Model(&Video{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))

	return true, nil
}

//删除用户点赞视频
func DeleteUserFavoriteVideo(user User, video_id int64) (bool, error) {

	if err := db.Where("user_id = ?  and video_id = ?", user.Id, video_id).Delete(&UserVideoFavorite{}).Error; err != nil {
		return false, err
	}

	db.Model(&Video{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))

	return true, nil
}

//点赞操作
func FavoriteAction(c *gin.Context) {

	token := c.Query("token")

	action_type := c.Query("action_type")

	video_id, _ := strconv.Atoi(c.Query("video_id"))

	if _, exist := usersLoginInfo[token]; exist {

		if action_type == "1" {
			SaveUserFavoriteVideo(usersLoginInfo[token], int64(video_id))
		} else {
			DeleteUserFavoriteVideo(usersLoginInfo[token], int64(video_id))
		}

		c.JSON(http.StatusOK, Response{StatusCode: 0})

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
//点赞视频列表
func FavoriteList(c *gin.Context) {

	token := c.Query("token")

	user, exist := usersLoginInfo[token]

	if exist {
		SearchUserFavorVideo(user)
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: FavoUserVideo,
	})

}

//查找点赞视频
func SearchUserFavorVideo(user User) (bool, error) {

	//先查找出在用户点赞过的视频id，再通过视频id找到视频。
	err := db.Where("id in (?)", db.Where("user_id = ?", user.Id).Select("video_id").Find(&FavoUserVideoList)).Find(&FavoUserVideo).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
