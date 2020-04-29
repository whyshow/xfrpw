package controllers

import (
	"../models"
	"../utils"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	"strconv"
	_ "strconv"
	_ "time"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) Home() {
	c.TplName = "index.html"
}

//web端根据方向获取课程
func (c *WebController) Classify() {
	code := c.GetString("c")
	o := orm.NewOrm()
	direction := []models.Ccit_classify{}
	o.QueryTable("ccit_classify").All(&direction)
	//查询所有分类
	for i := 0; i < len(direction); i++ {
		if code == direction[i].Indexes {
			c.Data["code"] = map[string]interface{}{"code": direction[i].Classify_name}
			ccit_course := []models.Ccit_course{}
			_, err := o.Raw("select * from ccit_course where classify_id = ?", direction[i].Classify_id).QueryRows(&ccit_course)
			if err == nil {
				// 查询成功
				c.Data["course"] = &ccit_course
			}
			break
		} else if code == "be" {
			ccit_course := []models.Ccit_course{}
			o.QueryTable("ccit_course").All(&ccit_course)
			c.Data["course"] = &ccit_course
			c.Data["code"] = map[string]interface{}{"code": "全部"}
		} else {
			//c.Abort("404")
			ccit_course := []models.Ccit_course{}
			o.QueryTable("ccit_course").All(&ccit_course)
			c.Data["course"] = &ccit_course
			c.Data["code"] = map[string]interface{}{"code": "全部"}
		}
	}
	c.Data["classify"] = &direction
	c.TplName = "html/classify.html"
}

//根据课程获取章节
func (c *WebController) Learn() {
	web_id := c.GetSession("web_user")
	if web_id == nil {
		c.Redirect("login", 302)
	} else {
		code := c.GetString("c")
		ccit_course := models.Ccit_course{}
		classify := models.Ccit_classify{}
		o := orm.NewOrm()
		ccit_course.Courset_id = code
		//查询课程表里的内容
		o.Read(&ccit_course)
		classify.Classify_id = ccit_course.Classify_id
		//查询方向表里面的内
		o.Read(&classify, "classify_id")

		var section []models.Ccit_section0
		var knobble []models.Ccit_knobble0
		o.Raw("SELECT * FROM ccit_section WHERE courset_id=" + "'" + code + "'" + " ORDER BY section_part ASC").QueryRows(&section)

		for index, _ := range section {
			sql := "SELECT * FROM ccit_knobble WHERE ccit_knobble.courset_id = " + "'" + code + "'" + " AND section_part = " + strconv.Itoa(index+1) + " ORDER BY ccit_knobble.knobble_name ASC"
			o.Raw(sql).QueryRows(&knobble)
			section[index].Ccit_knobble = knobble
		}

		c.Data["knobble"] = &knobble
		c.Data["classify"] = &classify
		c.Data["course"] = &ccit_course
		c.Data["section"] = &section
		c.TplName = "html/learn.html"
	}
}

//根据小节ID获得视频地址
func (c *WebController) Play() {
	o := orm.NewOrm()
	knobble := models.Ccit_knobble{}
	knobble.Knobble_id = c.GetString("c")
	o.Read(&knobble, "knobble_id")
	c.Data["knobble"] = &knobble
	c.TplName = "html/play.html"
}

func (c *WebController) Login() {
	if user := c.GetString("user"); user == "" {
		c.TplName = "html/login.html"
	} else {
		o := orm.NewOrm()
		ccit_user := []models.Ccit_user{}
		sql := "SELECT * FROM ccit_user WHERE student_number = " + "'" + user + "'" + "OR student_phone = " + "'" + user + "'" +
			" AND student_password = " + "'" + utils.MD5(c.GetString("password")) + "'"
		if num, _ := o.Raw(sql).QueryRows(&ccit_user); num > 0 {
			c.SetSession("web_user", user) //加入Session
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 2, "message": "账号或密码错误"}
			c.ServeJSON()
		}
	}
}
