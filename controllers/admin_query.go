package controllers

import (
	"../models"
	"../utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"strconv"
)

type DController struct {
	beego.Controller
}

//专业方向页面  查询所有的专业方向，然后将它发送到前端模板中用于展示。
func (c *DController) Get() {
	// 查询数据库 ccit_direction数据库，将结果发送到模板展示
	o := orm.NewOrm()
	page := c.GetString("page")
	direction := []models.Ccit_classify{}
	nums, _ := o.QueryTable("ccit_classify").All(&direction)
	if page == "" || page == "1" {
		sql := "SELECT * FROM ccit_classify limit 0,9"
		o.Raw(sql).QueryRows(&direction)
		c.Data["classify"] = &direction
		res := utils.Paginator(1, 9, nums)
		c.Data["paginator"] = res
		c.TplName = "admin/direction.html"
	} else {
		//将字符串类型转换成int类型
		p, _ := strconv.Atoi(page)
		res := utils.Paginator(p, 9, nums)
		s := p*9 - 9
		sql := "SELECT * FROM ccit_classify limit " + "'" + strconv.Itoa(s) + "'" + "," + "9"
		o.Raw(sql).QueryRows(&direction)
		c.Data["paginator"] = res
		c.Data["classify"] = &direction
		c.TplName = "admin/direction.html"

	}
}

type CourseController struct {
	beego.Controller
}

//课程管理页面   查询所有的课程，然后将它发送到前端模板中展示。
func (c *CourseController) Get() {
	name := c.GetString("Name")
	if name == "" {
		o := orm.NewOrm()
		ccit_course := []models.Ccit_course0{}
		direction := []models.Ccit_classify{}
		o.QueryTable("ccit_classify").All(&direction)

		sql := "SELECT ccit_course.*,ccit_classify.classify_name,ccit_classify.indexes FROM ccit_course, ccit_classify where  ccit_course.classify_id = ccit_classify.classify_id"
		nums, _ := o.Raw(sql).QueryRows(&ccit_course) //查询所有课程 和行数

		page := c.GetString("page")
		//当 page 为空或者为 1   则为第一页  否则不是第一页
		if page == "" || page == "1" {
			sql := "SELECT ccit_course.*,ccit_classify.classify_name,ccit_classify.indexes FROM ccit_course, ccit_classify where  ccit_course.classify_id = ccit_classify.classify_id ORDER BY putaway_time DESC limit 0,9"
			o.Raw(sql).QueryRows(&ccit_course)   //查询所有课程 和行数
			c.Data["ccit_course"] = &ccit_course //保存发送到前端
			res := utils.Paginator(1, 9, nums)   //制作分页
			c.Data["paginator"] = res            //将分页发送到前端
			c.Data["ccit_classify"] = direction  //将查询到的分类发送到前端
			c.TplName = "admin/course.html"
		} else {
			//将字符串类型转换成int类型
			p, _ := strconv.Atoi(page)         //获得当前点击的是第几页
			res := utils.Paginator(p, 9, nums) //制作分页
			s := p*9 - 9                       //计算从数据库第几行开始去数据
			sql := "SELECT * FROM ccit_course" + " " + "ORDER BY putaway_time DESC" + " " + "limit " + "'" + strconv.Itoa(s) + "'" + "," + "9"
			o.Raw(sql).QueryRows(&ccit_course) //查询
			c.Data["paginator"] = res
			c.Data["ccit_course"] = &ccit_course
			c.Data["ccit_classify"] = direction //将查询到的分类发送到前端
			c.TplName = "admin/course.html"
			return
		}

	} else {
		//下面的逻辑和上面基本一致，只不过在SQL查询上有所区别。

		o := orm.NewOrm()
		ccit_course := []models.Ccit_course{}
		page := c.GetString("page")
		//获得总的行数
		nums, _ := o.QueryTable("ccit_course").All(&ccit_course)
		//当 page 为空或者为 1   则为第一页  否则不是第一页
		if page == "" || page == "1" {
			sql := "SELECT * FROM ccit_course limit 0,9"
			o.Raw(sql).QueryRows(&ccit_course)
			c.Data["ccit_course"] = &ccit_course
			res := utils.Paginator(1, 9, nums)
			c.Data["paginator"] = res
			c.TplName = "admin/course.html"
		} else {
			//将字符串类型转换成int类型
			p, _ := strconv.Atoi(page)         //获得当前点击的是第几页
			res := utils.Paginator(p, 9, nums) //制作分页
			s := p*9 - 9                       //计算从数据库第几行开始去数据
			sql := "SELECT * FROM ccit_course limit " + "'" + strconv.Itoa(s) + "'" + "," + "9"
			o.Raw(sql).QueryRows(&ccit_course) //查询
			c.Data["paginator"] = res
			c.Data["ccit_course"] = &ccit_course
			c.TplName = "admin/course.html"
			return
		}
		return
	}
}

type DirectionController struct {
	beego.Controller
}

//通过Ajax 获得单个方向的所有课程
func (c *DirectionController) Post() {
	classify_id := c.GetString("id") //获得方向id
	ccit_course := []models.Ccit_course0{}
	o := orm.NewOrm()
	sql := "SELECT ccit_course.*,ccit_classify.classify_name,ccit_classify.indexes FROM ccit_course, ccit_classify where  ccit_course.classify_id = ccit_classify.classify_id and " + "ccit_course.classify_id = " + "'" + classify_id + "'"
	nums, _ := o.Raw(sql).QueryRows(&ccit_course) //查询所有课程 和行数
	page := c.GetString("page")
	if page == "" || page == "1" {
		sql := "SELECT ccit_course.*,ccit_classify.classify_name,ccit_classify.indexes FROM ccit_course, ccit_classify where  ccit_course.classify_id = ccit_classify.classify_id and ccit_course.classify_id = " + "'" + classify_id + "'" + "ORDER BY putaway_time DESC limit 0,9"
		o.Raw(sql).QueryRows(&ccit_course)   //查询所有课程 和行数
		c.Data["ccit_course"] = &ccit_course //保存发送到前端
		res := utils.Paginator(1, 9, nums)   //制作分页
		c.Data["paginator"] = res            //将分页发送到前端
		c.Data["json"] = map[string]interface{}{"code": 1, "message": &ccit_course, "page": &res}
		c.ServeJSON()
	} else {

	}
}

type SectionController struct {
	beego.Controller
}

//章节管理页面   查询所有的章节，然后将它发送到前端模板中展示。
func (c *SectionController) Get() {
	o := orm.NewOrm()
	classify := []models.Ccit_classify{}
	sql := "SELECT classify_id,classify_name FROM `ccit_classify`"
	o.Raw(sql).QueryRows(&classify)
	c.Data["ccit_course"] = &classify
	c.TplName = "admin/section.html"

}

type GetSectionController struct {
	beego.Controller
}

//根据课程id查询课程内的章节信息      得到ID ->向数据库查询 ->发送到前端展示
func (c *GetSectionController) Post() {
	//得到ID
	Courset_id := c.GetString("Courset_id")
	if len(Courset_id) < 5 {
		c.Data["json"] = map[string]interface{}{"code": -2, "message": "查询出错"}
		c.ServeJSON()
	} else {
		sql := "SELECT * FROM ccit_section WHERE courset_id=" + "'" + Courset_id + "'"
		o := orm.NewOrm()
		//切片
		section := []models.Ccit_section{}
		//查询
		if num, err := o.Raw(sql).QueryRows(&section); err != nil {
			//打印错误日志
			c.Data["json"] = map[string]interface{}{"code": 2, "message": "查询出错"}
			c.ServeJSON()
		} else if num > 0 {
			//查询到数据
			c.Data["json"] = map[string]interface{}{"code": 1, "message": &section}
			c.ServeJSON()
		} else {
			//未查询到任何数据
			c.Data["json"] = map[string]interface{}{"code": 2, "message": "没有数据"}
			c.ServeJSON()
		}

	}
}

type GetCourseController struct {
	beego.Controller
}

//根据方向内容查询课程列表
func (c *GetCourseController) Post() {
	//得到专业方向id
	classify_id := c.GetString("classify_id")
	course := []models.Ccit_course{}
	sql := "SELECT courset_id,course_name FROM `ccit_course` WHERE classify_id =" + "'" + classify_id + "'"
	o := orm.NewOrm()
	r, _ := o.Raw(sql).QueryRows(&course)
	if r == 0 {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "暂无课程"}
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": &course}
		c.ServeJSON()
		return
	}

}

type LogsController struct {
	beego.Controller
}

func (c *LogsController) Get() {
	//打开文件
	//f, err := os.Open("logs/info.log")
	//if err != nil {
	//	return
	//}
	////关闭流
	//defer f.Close()
	////读取文件内的信息
	//scanner := bufio.NewScanner(f)
	//var s  =""
	////info := Logs_info{}
	////循环读取每行信息
	//for scanner.Scan()  {
	//	s = s+"{"+"Txt:"+scanner.Text()+"}"
	//}
	//c.Data["info"] ="["+s+"]"
	//fmt.Println("["+s+"]")
	//c.TplName = "html/logs.html"
	b, err := ioutil.ReadFile("logs/admin.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Data["info"] = string(b)
	c.TplName = "admin/info_log.html"
	//下载
	//c.Ctx.Output.Download("logs/info.txt","info.txt")
}
