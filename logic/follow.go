package logic

import (
	"Tiktok/model"
)

type FollowServiceImp struct {
	model.FollowService
	model.UserService
}

/*
	对外提供服务之返回登陆用户的关注用户数量
*/

// GetFollowingCnt 加入redis 根据用户id查询关注数
func (followService *FollowServiceImp) GetFollowingCnt(userId int64) (int64, error) {
	//followDao := mysql.NewFollowDaoInstance()
	return 0, nil
}
