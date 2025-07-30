package api

import (
	"Tiktok/dao/mysql"
	"Tiktok/logic"
	"Tiktok/model"
	"Tiktok/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user-id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	userId := c.Query("id")
	// 使用中间件，做token权限校验
	usi := logic.GetUserServiceInstance()
	id, _ := strconv.ParseInt(userId, 10, 64)
	if user, err := usi.GetUserLoginInfoById(id); err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Query Success"},
			User:     user,
		})
	}
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	usi := logic.GetUserServiceInstance()
	user := usi.GetUserBasicInfoByName(username)
	if username == user.Name {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		newUser := mysql.UserBasicInfo{
			Name:     username,
			Password: logic.Encoder(password),
		}
		if usi.InsertUser(&newUser) != true {
			fmt.Println("Insert Fail")
		}
		// 得到用户id
		user := usi.GetUserBasicInfoByName(username)
		userId := user.Id
		token := utils.GenerateToken(userId, username)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "Register Success",
			},
			UserId: userId,
			Token:  token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	encoderPassword := logic.Encoder(password)
	usi := logic.GetUserServiceInstance()
	user := usi.GetUserBasicInfoByName(username)
	userId := user.Id
	if encoderPassword == user.Password {
		token := utils.GenerateToken(userId, username)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Login Success"},
			UserId:   userId,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User or Password Error",
			},
		})
	}
}
