package user

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type Controller struct {
	web.Controller
}

func (this *Controller) Get() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err == nil && id > 0 {
		o := orm.NewOrm()
		user := User{Id: id}
		err = o.Read(&user)
		if err != nil {
			if err == orm.ErrNoRows {
				this.CustomAbort(404, "User not found")
			}
			this.CustomAbort(500, err.Error())
		}
		this.Data["json"] = user
	}

	this.ServeJSON()
}

func (this *Controller) Post() {
	var user User
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err != nil {
		this.CustomAbort(400, err.Error())
	}
	o := orm.NewOrm()
	user.Id = 0
	if _, err := o.Insert(&user); err != nil {
		this.CustomAbort(404, err.Error())
	}
	this.Data["json"] = user
	this.ServeJSON()
}

func (this *Controller) Put() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil && id <= 0 {
		this.CustomAbort(404, "User not found")
	}
	o := orm.NewOrm()
	user := User{Id: id}
	if o.Read(&user) != nil {
		this.CustomAbort(404, "User not found")
	}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err != nil {
		this.CustomAbort(400, err.Error())
	}
	user.Id = id
	if _, err = o.Update(&user); err != nil {
		this.CustomAbort(400, err.Error())
	}
	this.Data["json"] = user
	this.ServeJSON()
}

func (this *Controller) Delete() {
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil && id <= 0 {
		this.CustomAbort(404, "User not found")
	}
	o := orm.NewOrm()
	this.Data["json"] = 0
	if num, err := o.Delete(&User{Id: id}); err == nil {
		this.Data["json"] = num
	}
	this.ServeJSON()
}
