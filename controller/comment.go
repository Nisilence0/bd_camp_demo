package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

//储存评论
func SaveComment(user User, video_id int64, text, formatTimeStr string) (bool, error) {

	if err := db.Create(&Comment{UserId: user.Id, VideoId: video_id, Content: text, CreateDate: formatTimeStr}).Error; err != nil {
		return false, err
	}

	db.Model(&Video{}).Where("id = ?", video_id).Update("comment_count", gorm.Expr("comment_count + ?", 1))

	return true, nil
}

//查找不同视频的评论
func SearchComment(video_id string) (bool, error) {

	if err := db.Preload(clause.Associations).Where("video_id = ?", video_id).Find(&DemoComments).Error; err != nil {
		return false, err
	}

	return true, nil

}

// CommentAction no practical effect, just check if token is valid
//评论行为
func CommentAction(c *gin.Context) {

	token := c.Query("token")
	actionType := c.Query("action_type")
	video_id, _ := strconv.Atoi(c.Query("video_id"))
	text := c.Query("comment_text")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {

			timeUnix := time.Now().Unix() //已知的时间戳
			formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
			SaveComment(usersLoginInfo[token], (int64)(video_id), text, formatTimeStr)

			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: formatTimeStr,
				}})
			return
		} else {

		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
//评论列表
func CommentList(c *gin.Context) {
	video_id := c.Query("video_id")

	SearchComment(video_id)

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
