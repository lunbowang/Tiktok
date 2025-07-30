package model

type FollowService interface {

	/*
		模块对外提供的服务接口
	*/
	// GetFollowingCnt 根据用户id查询关注数
	GetFollowingCnt(userId int64) (int64, error)
}
