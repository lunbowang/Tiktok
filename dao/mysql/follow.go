package mysql

import "sync"

type Follow struct {
	Id          int64
	UserId      int64
	FollowingId int64
	Followed    int8
	CreatedAt   string
	UpdatedAt   string
}

func (Follow) TableName() string {
	return "relation"
}

type FollowDao struct {
}

var (
	followDao  *FollowDao
	followOnce sync.Once
)

// NewFollowDaoInstance 生成并返回followDao的单例对象。
func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		},
	)
	return followDao
}
