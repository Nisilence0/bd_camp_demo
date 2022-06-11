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

func AddFollowRelat(user, author int64) (bool, error) {

	if err := db.Create(&UserUserFollow{Id: author, FansId: user}).Error; err != nil {
		return false, err
	}

	db.Model(&User{}).Where("id = ?", author).Update("follower_count", gorm.Expr("follower_count + ?", 1))
	db.Model(&User{}).Where("id = ?", user).Update("follow_count", gorm.Expr("follow_count + ?", 1))

	return true, nil

}

func RemoveFollowRelat(user, author int64) (bool, error) {

	if err := db.Create(&UserUserFollow{Id: author, FansId: user}).Error; err != nil {
		return false, err
	}

	db.Model(&User{}).Where("id = ?", author).Update("follower_count", gorm.Expr("follower_count - ?", 1))
	db.Model(&User{}).Where("id = ?", user).Update("follow_count", gorm.Expr("follow_count - ?", 1))

	return true, nil

}

func SelectFollower(user_Id int64) (bool, error) {

	if err := db.Where("id in (?)", db.Where("fans_id = ?", user_Id).Find(&UserFollowerIdList)).Find(&UserFollowerList).Error; err != nil {

		return false, err
	}
	return true, nil
}

func SelectFollow(user_Id int64) (bool, error) {

	if err := db.Where("id in (?)", db.Where("id = ?", user_Id).Find(&UserFollowIdList)).Find(&UserFollowList).Error; err != nil {

		return false, err
	}
	return true, nil
}

// RelationAction no practical effect, just check if token is valid
// 没有实用效果，只检查token令牌是否合法
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	action_type := c.Query("action_type")

	to_user_id, _ := strconv.Atoi(c.Query("to_user_id"))

	if _, exist := usersLoginInfo[token]; exist {
		user := usersLoginInfo[token]
		if action_type == "1" {
			AddFollowRelat(user.Id, (int64)(to_user_id))
		} else {
			RemoveFollowRelat(user.Id, (int64)(to_user_id))
		}

		c.JSON(http.StatusOK, Response{StatusCode: 0})

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't login"})
	}
}

// FollowList all users have same follow list
// 每个粉丝列表都相同
func FollowList(c *gin.Context) {

	user_id, _ := (strconv.Atoi(c.Query("user_id")))

	SelectFollower((int64)(user_id))

	fmt.Printf("%v\n", &UserFollowIdList)

	fmt.Printf("%v\n", &UserFollowList)

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: UserFollowerList,
	})
}

// FollowerList all users have same follower list
// 每个粉丝的关注列表都相同
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
