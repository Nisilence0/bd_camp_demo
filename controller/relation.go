package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

//清除当前所有用户的关注
func ClearUserFollow() (bool, error) {

	if err := db.Model(&User{}).Where("is_follow = ?", 1).Update("is_follow", 0).Error; err != nil {
		return false, err
	}
	return true, nil
}

//添加关注关系
func AddFollowRelat(user, author int64) (bool, error) {

	if err := db.Create(&UserUserFollow{Id: author, FansId: user}).Error; err != nil {
		return false, err
	}

	db.Model(&User{}).Where("id = ?", author).Update("follower_count", gorm.Expr("follower_count + ?", 1))
	db.Model(&User{}).Where("id = ?", user).Update("follow_count", gorm.Expr("follow_count + ?", 1))

	db.Model(&User{}).Where("id = ?", author).Update("is_follow", 1)

	return true, nil

}

//移除关注关系
func RemoveFollowRelat(user, author int64) (bool, error) {

	if err := db.Delete(&UserUserFollow{Id: author, FansId: user}).Error; err != nil {
		return false, err
	}

	db.Model(&User{}).Where("id = ?", author).Update("follower_count", gorm.Expr("follower_count - ?", 1))
	db.Model(&User{}).Where("id = ?", user).Update("follow_count", gorm.Expr("follow_count - ?", 1))

	db.Model(&User{}).Where("id = ?", author).Update("is_follow", 0)
	return true, nil

}

//查找粉丝
func SelectFollower(user_Id int64) (bool, error) {

	if err := db.Where("id in (?)", db.Where("fans_id = ?", user_Id).Select("id").Find(&UserFollowerIdList)).Find(&UserFollowerList).Error; err != nil {

		return false, err
	}
	return true, nil
}

//查找关注用户
func SelectFollow(user_Id int64) (bool, error) {

	if err := db.Where("id in (?)", db.Where("id = ?", user_Id).Select("fans_id").Find(&UserFollowIdList)).Find(&UserFollowList).Error; err != nil {

		return false, err
	}
	return true, nil
}

//关注行为返回
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	action_type := c.Query("action_type")

	to_user_id, _ := strconv.Atoi(c.Query("to_user_id"))

	if _, exist := usersLoginInfo[token]; exist {
		user := usersLoginInfo[token]
		if action_type == "1" {
			AddFollowRelat(user.Id, (int64)(to_user_id))
		} else {
			fmt.Println(78)
			RemoveFollowRelat(user.Id, (int64)(to_user_id))
		}

		c.JSON(http.StatusOK, Response{StatusCode: 0})

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't login"})
	}
}

//关注列表
func FollowList(c *gin.Context) {

	user_id, _ := (strconv.Atoi(c.Query("user_id")))

	SelectFollower((int64)(user_id))

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: UserFollowerList,
	})
}

//粉丝列表
func FollowerList(c *gin.Context) {

	user_id, _ := (strconv.Atoi(c.Query("user_id")))

	SelectFollow((int64)(user_id))

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: UserFollowList,
	})
}
