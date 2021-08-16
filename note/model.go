package note

type Note struct {
	Id   int
	Data string `orm:"size(100)"`
	User int
}
