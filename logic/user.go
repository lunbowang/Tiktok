package logic

import (
	"Tiktok/config"
	"Tiktok/dao/mysql"
	"Tiktok/model"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"sync"
)

type UserServiceImpl struct {
	// 关注服务
	model.FollowService
}

var (
	userServiceImp  *UserServiceImpl
	userServiceOnce sync.Once
)

func GetUserServiceInstance() *UserServiceImpl {
	userServiceOnce.Do(func() {
		userServiceImp = &UserServiceImpl{
			FollowService: &FollowServiceImp{},
		}
	})
	return userServiceImp
}

func (usi *UserServiceImpl) GetUserBasicInfoByName(name string) mysql.UserBasicInfo {
	user, err := mysql.GetUserBasicInfoName(name)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return user
	}
	log.Println("Query User Success")
	return user
}

func (usi *UserServiceImpl) InsertUser(user *mysql.UserBasicInfo) bool {
	flag := mysql.InsertUser(user)
	if flag == false {
		log.Println("Insert Fail!")
		return false
	}
	return true
}

// GetUserLoginInfoById 未登录情况返回用户信息
func (usi *UserServiceImpl) GetUserLoginInfoById(id int64) (model.User, error) {
	user := model.User{
		Id:              5,
		Name:            "qcj",
		FollowCount:     1,
		FollowerCount:   99999,
		IsFollow:        false,
		Avatar:          config.CUSTOM_DOMAIN + config.OSS_USER_AVATAR_DIR,
		BackgroundImage: config.BG_IMAGE,
		Signature:       config.SIGNATURE,
		TotalFavorited:  10,
		FavoriteCount:   10,
		WorkCount:       8,
	}
	u, err := mysql.GetUserBasicInfoById(id)
	fmt.Println(u)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
	}
	user.Id = u.Id
	user.Name = u.Name
	//userService := GetUserServiceInstance()
	//var wg sync.WaitGroup
	//wg.Add(5)
	// todo 计算关注数

	// todo 计算粉丝数

	// todo 计算作品数

	// todo 计算被点赞数, 找出用户被点赞的视频，循环求和:在likeservide实现

	// todo 计算喜欢数

	return user, nil
}

// Encoder 加密
func Encoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
