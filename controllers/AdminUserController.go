package controllers

import (
	"../models"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
)

type AdminUserController struct {
	beego.Controller
}

//请求登录界面
func (c *AdminUserController) LoginPage() {
	c.TplName = "admin/login_page.html"
}

//登录请求
func (c *AdminUserController) Login() {
	// 1.接收数据
	user := c.GetString("Name")
	password := c.GetString("Password")
	// 2.验证账户密码是否正确
	ccitUser := models.Ccit_admin{}
	ccitUser.Admin_account = user
	ccitUser.Admin_password = password
	if code := models.AdminLogin(ccitUser, password, c.Ctx.Request); code == 1 {
		//登录失败 用户名不存在
		c.Data["json"] = map[string]interface{}{"code": code, "message": "用户名不存在"}
		c.ServeJSON()
	} else if code == 2 {
		//登录失败 密码错误
		c.Data["json"] = map[string]interface{}{"code": code, "message": "密码错误"}
		c.ServeJSON()
	} else if code == 0 {
		//登录成功 加入session
		c.SetSession("admin_id", ccitUser.Admin_account)
		c.Data["json"] = map[string]interface{}{"code": code, "message": "/admin/index"}
		c.ServeJSON()
	}
}

//退出请求
func (c *AdminUserController) Exit() {
	models.AdminiExit(c.GetSession("admin_id").(string), c.Ctx.Request)
	c.DelSession("admin_id")
	c.TplName = "admin/login_page.html"
}

//后台管理首页
func (c *AdminUserController) Home() {
	c.TplName = "admin/index.html"
}

//数据统计页面 or 请求
func (c *AdminUserController) Statistics() {
	//暂时还没写
	c.TplName = "admin/statisticsView.html"
}
