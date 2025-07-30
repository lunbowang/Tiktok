package setting

import "Tiktok/dao/mysql"

type database struct {
}

func (db database) Init() {
	mysql.Init()
}
