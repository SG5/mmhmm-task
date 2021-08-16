package main

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"mmhmm-task/note"
	"mmhmm-task/user"
)

func init() {
	// register model
	orm.RegisterModel(new(user.User))
	orm.RegisterModel(new(note.Note))

	// set default database
	orm.RegisterDataBase("default", "sqlite3", "file::?cache=shared&mode=memory")

	o := orm.NewOrm()
	o.Raw("CREATE TABLE user (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)").Exec()
	o.Raw("CREATE TABLE note (id INTEGER PRIMARY KEY AUTOINCREMENT, data TEXT, user INTEGER)").Exec()
}
