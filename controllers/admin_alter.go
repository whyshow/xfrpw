package controllers

import (
	"../models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type AlterFxPageController struct {
	beego.Controller
}

//修改方向的页面
func (c *AlterFxPageController) Get() {
	Id := c.GetString("id") //得到 id
	o := orm.NewOrm()
	classify := []models.Ccit_classify{}
	sql := "SELECT * FROM ccit_classify WHERE classify_id =" + Id
	o.Raw(sql).QueryRows(&classify)
	c.Data["data"] = &classify
	c.TplName = "admin/alterzyfx.html"

}

type AlterFxController struct {
	beego.Controller
}

//修改专业方向的操作
func (c *AlterFxController) Get() {
	Index := c.GetString("Index") //得到 index
	Name := c.GetString("Name")   //得到 name
	Id := c.GetString("id")       //得到 id
	o := orm.NewOrm()
	if Index != "" && Name != "" && Id != "" {
		if _, err := o.Update(&models.Ccit_classify{Classify_id: Id, Classify_name: Name, Add_date: time.Now().Format("2006-01-02 15:04:05"),
			Indexes: Index}); err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "修改成功"} //制作json数据返回给Ajax
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改失败"} //制作json数据返回给Ajax
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "数据不能为空"} //制作json数据返回给Ajax
		c.ServeJSON()
	}

}

type AlterKcPageController struct {
	beego.Controller
}

//修改课程的页面
func (c *AlterKcPageController) Get() {
	o := orm.NewOrm()
	course := models.Ccit_course{}
	classify := models.Ccit_classify{}
	course.Courset_id = c.GetString("id")

	err := o.Read(&course, "courset_id")
	if err == nil {
		classify.Classify_id = course.Classify_id
		o.Read(&classify, "classify_id")
		c.Data["data"] = &course
		c.Data["classify"] = &classify
		c.TplName = "admin/alterkc.html"
	} else {

	}
}

type AlterKcController struct {
	beego.Controller
}

func (c *AlterKcController) Get() {

}
