package note

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"mmhmm-task/user"
	"strconv"
)

type Controller struct {
	web.Controller
}

func (this *Controller) Get() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err == nil && id > 0 {
		o := orm.NewOrm()
		Note := Note{Id: id}
		err = o.Read(&Note)
		if err != nil {
			if err == orm.ErrNoRows {
				this.CustomAbort(404, "Note not found")
			}
			this.CustomAbort(500, err.Error())
		}
		this.Data["json"] = Note
	}

	this.ServeJSON()
}

func (this *Controller) Post() {
	var note Note
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &note); err != nil {
		this.CustomAbort(400, err.Error())
	}
	o := orm.NewOrm()
	user := user.User{Id: note.User}
	if o.Read(&user) != nil {
		this.CustomAbort(404, "User not found")
	}
	note.Id = 0
	if _, err := o.Insert(&note); err != nil {
		this.CustomAbort(404, err.Error())
	}
	this.Data["json"] = note
	this.ServeJSON()
}

func (this *Controller) Put() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil && id <= 0 {
		this.CustomAbort(404, "Note not found")
	}
	o := orm.NewOrm()
	note := Note{Id: id}
	if o.Read(&note) != nil {
		this.CustomAbort(404, "Note not found")
	}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &note); err != nil {
		this.CustomAbort(400, err.Error())
	}
	user := user.User{Id: note.User}
	if o.Read(&user) != nil {
		this.CustomAbort(404, "User not found")
	}
	note.Id = id
	if _, err = o.Update(&note); err != nil {
		this.CustomAbort(400, err.Error())
	}
	this.Data["json"] = note
	this.ServeJSON()
}

func (this *Controller) Delete() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil && id <= 0 {
		this.CustomAbort(404, "Note not found")
	}
	o := orm.NewOrm()
	this.Data["json"] = 0
	if num, err := o.Delete(&Note{Id: id}); err == nil {
		this.Data["json"] = num
	}
	this.ServeJSON()
}
