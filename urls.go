package main

import (
	"github.com/beego/beego/v2/server/web"
	"mmhmm-task/note"
	"mmhmm-task/user"
)

func init() {
	web.Router("/user/:id([0-9]+", &user.Controller{})
	web.Router("/user", &user.Controller{})
	web.Router("/note/:id([0-9]+", &note.Controller{})
	web.Router("/note", &note.Controller{})
}
