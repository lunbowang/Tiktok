package mysql

import (
	"log"
	"time"
)

type UserBasicInfo struct {
	Id       int64
	Name     string
	Password string
	CreateAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (user UserBasicInfo) TableName() string {
	return "user"
}

func GetUserBasicInfoName(name string) (UserBasicInfo, error) {
	user := UserBasicInfo{}
	// 使用 GORM 查询数据库，通过用户名精准匹配首条记录
	// 注意：若存在重名用户，仅返回数据库中排序第一的记录
	if err := Db.Where("name = ?", name).First(&user).Error; err != nil {
		log.Println("获取用户信息读库失败", err.Error())
		return user, err
	}
	return user, nil
}

func InsertUser(user *UserBasicInfo) bool {
	if err := Db.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func GetUserBasicInfoById(id int64) (UserBasicInfo, error) {
	user := UserBasicInfo{}
	if err := Db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}
