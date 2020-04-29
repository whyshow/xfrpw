package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type AndroidController struct {
	beego.Controller
}

//移动端查询专业方向信息
func (c *AndroidController) Classify() {
	//初始化数据库连接池
	o := orm.NewOrm()
	//得到结构图切片
	ccit_classify := []models.Ccit_classify{}
	num, err := o.Raw("select * from ccit_classify").QueryRows(&ccit_classify)
	//判断查询时是否出现错误
	if err == nil {
		//判断查询结果是否有数据
		if num > 0 {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询成功", "result": &ccit_classify}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "没有专业方向", "result": nil}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": err, "result": nil}
		c.ServeJSON()
	}
}

//移动端查询课程信息
func (c *AndroidController) Course() {
	course := models.Ccit_course{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &course)
	//判断参数 code是否为空
	if course.Classify_id != "" {
		o := orm.NewOrm()
		ccit_course := []models.Ccit_course{}
		num, err := o.Raw("select * from ccit_course where classify_id = ?", course.Classify_id).QueryRows(&ccit_course)
		//判断查询时是否出现错误
		if err == nil {
			//判断查询结果是否有数据
			if num > 0 {
				c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询成功", "result": &ccit_course}
				c.ServeJSON()
			} else {
				c.Data["json"] = map[string]interface{}{"code": 1, "message": "暂无课程", "result": nil}
				c.ServeJSON()
			}
		} else {
			c.Data["json"] = map[string]interface{}{"code": -1, "message": err, "result": nil}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "参数错误", "result": nil}
		c.ServeJSON()
	}
}

//移动端查询章节信息
func (c *AndroidController) Learn() {
	//得到课程表结构体
	course := models.Ccit_course{}
	//将post的json数据映射到课程表结构体中
	json.Unmarshal(c.Ctx.Input.RequestBody, &course)
	if course.Courset_id != "" {
		o := orm.NewOrm()
		//课程表结构体
		ccit_course := models.Ccit_course{}
		//专业方向表结构体
		classify := models.Ccit_classify{}
		//将课程ID赋值给课程表结构体
		ccit_course.Courset_id = course.Courset_id
		//查询课程表里的内容 根据课程ID查询到专业方向ID
		o.Read(&ccit_course)
		//将得到的专业方向ID赋值给专业方向结构体
		classify.Classify_id = ccit_course.Classify_id
		//查询专业方向表里面的内容
		o.Read(&classify, "classify_id")
		//章节表结构体
		var section []models.Ccit_section0
		//小节表结构体
		var knobble []models.Ccit_knobble0
		//根据课程ID查询章节
		o.Raw("SELECT * FROM ccit_section WHERE courset_id=" + "'" + course.Courset_id + "'" + " ORDER BY section_part ASC").QueryRows(&section)
		for index, _ := range section {
			sql := "SELECT * FROM ccit_knobble WHERE ccit_knobble.courset_id = " + "'" + course.Courset_id + "'" + " AND section_part = " + strconv.Itoa(index+1) + " ORDER BY ccit_knobble.knobble_name ASC"
			o.Raw(sql).QueryRows(&knobble)
			section[index].Ccit_knobble = knobble
		}
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询成功", "result": map[string]interface{}{"classify": &classify, "course": &ccit_course, "section": &section}}
		c.ServeJSON()

	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "参数错误", "result": nil}
		c.ServeJSON()
	}
}

//移动端查询视频地址
func (c *AndroidController) Play() {
	knobble := models.Ccit_knobble{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &knobble)
	if knobble.Knobble_id != "" {
		o := orm.NewOrm()
		o.Read(&knobble, "knobble_id")
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "查询成功", "result": knobble}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "参数错误", "result": nil}
		c.ServeJSON()
	}
}

//移动端用户登录
func (c *AndroidController) Login() {
	//得到课程表结构体
	user := models.Ccit_user{}
	//将post的json数据映射到课程表结构体中
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	user.Student_password = utils.MD5(user.Student_password)
	if ok, code, message, user := models.Login(user); ok {
		c.Data["json"] = map[string]interface{}{"code": code, "message": message, "result": &user}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": code, "message": message, "result": nil}
		c.ServeJSON()
	}
}

//移动端找回密码
func (c *AndroidController) Findpassword() {
	user := models.Ccit_user{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	user.Student_password = utils.MD5(user.Student_password)
	newUser := models.Ccit_user{Student_number: user.Student_number, Student_phone: user.Student_phone}
	if ok, message := models.FindUserPassword(newUser); ok {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": message, "result": nil}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": message, "result": nil}
		c.ServeJSON()
	}
}
