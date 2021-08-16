package user

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}
