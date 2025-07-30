package setting

type group struct {
	DB database
}

var Group = new(group)

func Init() {
	Group.DB.Init()
}
