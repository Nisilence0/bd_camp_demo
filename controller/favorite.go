package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
// FavoriteAction无实际效果，只需检查令牌是否有效

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
// Favoritelist 所有用户都有相同的喜爱视频列表
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
