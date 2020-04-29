package controllers

import (
	"../models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) QueryUser() {
	user := models.Ccit_user{Student_number: c.GetString("Student_number")}
	if ok, _, _ := models.QueryUserInfo(user); ok {

	} else {

	}
}

func (c *UserController) QueryUsers() {
	if ok, _, users := models.QueryUserInfos(); ok {
		c.Data["user"] = &users
		c.TplName = "admin/user.html"
	} else {
		c.TplName = "admin/user.html"
	}
}

func (c *UserController) AddUserPage() {
	c.TplName = "admin/tj-users.html"
}

func (c *UserController) AddUsers() {

}
