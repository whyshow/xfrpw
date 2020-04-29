package controllers

import (
	"../models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)

type RemoveFController struct {
	beego.Controller
}

//专业方向 删除操作
func (c *RemoveFController) Get() {
	//得到 专业方向的 id
	Id := c.GetString("id")
	o := orm.NewOrm()
	//得到课程数据库模型切片
	course := []models.Ccit_course{}
	//这个SQL语句是为了查询专业方向所属的课程，如果查询到了课程那么这个专业方向就不能删除
	sql := "SELECT * FROM ccit_course WHERE classify_id =" + "'" + Id + "'"
	//查询数据库数据， _ 返回的错误信息、num 返回查询到的行数。
	num, _ := o.Raw(sql).QueryRows(&course)
	if num == 0 { //没有查询到，可以进行删除了     ps: Delete(切片)
		if _, err := o.Delete(&models.Ccit_classify{Classify_id: Id}); err == nil {
			if os.RemoveAll("static/upvideo/" + Id); err != nil {
				//删除失败  错误日志

			}
			//制作json数据返回给Ajax
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
			c.ServeJSON()
		}
	} else { // ajax通知前端 这个分类方向内有课程
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "方向内有课程"}
		c.ServeJSON()
	}
}

type RemoveKCController struct {
	beego.Controller
}

//删除课程 操作
func (c *RemoveKCController) Get() {
	//得到 要删除课程的·id
	Id := c.GetString("id")
	o := orm.NewOrm()
	//得到章节数据库模型切片
	section := []models.Ccit_section{}
	//这个SQL语句是为了查询课程所属的章节，如果查询到了章节信息那么这个课程就不能删除
	sql := "SELECT * FROM ccit_section WHERE courset_id =" + "'" + Id + "'"
	//查询数据库数据， _ 返回的错误信息、num 返回查询到的行数。
	num, _ := o.Raw(sql).QueryRows(&section)
	if num == 0 { //没有查询到，可以进行删除了     ps: Delete(切片)
		if _, err := o.Delete(&models.Ccit_course{Courset_id: Id}); err == nil {
			//制作json数据返回给Ajax
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
			c.ServeJSON()
		}
	} else { // ajax通知前端 这个课程内有章节
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "课程内有章节数据"}
		c.ServeJSON()
	}
}
