package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/pkg/util"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

//Create a new User and save User information
//创建一个新的用户，并保存用户信息
func CreateUser(username, password string) (int64, error) {

	if err := db.Create(&UserLogin{Name: username, Password: password}).Error; err != nil {
		return -1, err
	}

	if err := db.Create(&User{Name: username}).Error; err != nil {
		return -1, err
	}
	var user User

	db.Where(&User{Name: username}).First(&user)

	return user.Id, nil
}

//查看用户是存在
func CheckUser(username, password string) (int64, error) {

	var userlogin UserLogin
	err := db.Where(&UserLogin{Name: username, Password: password}).First(&userlogin).Error
	if err != nil {
		return -1, err
	}

	return userlogin.Id, nil
}

//获取用户信息
func GetUser(name, token string, id int64) (User, error) {
	var user User
	err := db.Where(&User{Id: id, Name: name}).First(&user).Error
	if err != nil {
		return user, err
	}
	usersLoginInfo[token] = user

	return user, nil
}

//用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := util.EncodeMD5(c.Query("password"))
	var token string
	var id int64

	_, err := CheckUser(username, password)
	if err != nil {
		id, _ = CreateUser(username, password)
		token = username + "&" + password
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User has exist"},
		})
		return
	}

	_, err = GetUser(username, token, id)

	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User regist error"},
		})
		return
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
	}

	GetAllVideos(c)
}

//用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := util.EncodeMD5(c.Query("password"))
	token := username + "&" + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
		return
	}

	id, err := CheckUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Username or Password is error"},
		})
		return
	}

	_, err = GetUser(username, token, id)

	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User logining error"},
		})
		return
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
	}

	GetAllVideos(c)
}

//用户信息返回
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
