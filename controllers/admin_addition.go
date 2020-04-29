package controllers

import (
	"../models"
	"../utils"
	"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"os"
	"time"
)

type TjzyfxController struct {
	beego.Controller
}

//1.添加专业方向页面 or 请求
//2.如果 GET的为空值则会返回页面，否在会向数据库添加信息
func (c *TjzyfxController) Get() {
	name := c.GetString("Name")
	index := c.GetString("Index")
	if name == "" {
		c.TplName = "admin/tjzyfx.html"
		return
	} else {
		// 以时间为信息生成随机数 此随机数 == Direction_id
		o := orm.NewOrm()
		ccit_direction := models.Ccit_classify{}
		ccit_direction.Classify_name = name                                //赋值 方向名称
		ccit_direction.Indexes = index                                     //赋值 索引
		ccit_direction.Classify_id = utils.Random()                        //生成ID,同时也作为存放方向内的课程视频的文件夹
		ccit_direction.Classify_path = ccit_direction.Classify_id          //专业方向文件夹名，和专业方向ID保持一致
		ccit_direction.Add_date = time.Now().Format("2006-01-02 15:04:05") //获取当前日期时间
		//创建一个这个专业方向的视频文件夹用于存放当前方向内所有课程的视频。
		os.Mkdir("static/upvideo/"+ccit_direction.Classify_path, os.ModePerm)
		//然后向创建的文件夹内放个文件，主要是描述一下这个文件夹是放什么的
		f, err1 := os.Create("static/upvideo/" + ccit_direction.Classify_path + "/README.md")
		defer f.Close()
		if err1 != nil {
			//错误日志
		} else {
			f.Write([]byte("这个文件夹存放的是 " + ccit_direction.Classify_name + " 方向内的课程"))
		}
		//向数据库插入数据
		_, err := o.Insert(&ccit_direction)
		if err == nil {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "/admin/direction.html"}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{"code": 2, "message": "方向已存在"}
			c.ServeJSON()
			return
		}
	}
}

type TjkcController struct {
	beego.Controller
}

//添加课程页面   将所有的课程查询出来发送到前端模板中用于下拉列表的展示。
func (c *TjkcController) Get() {
	o := orm.NewOrm()
	direction := []models.Ccit_classify{}
	o.QueryTable("ccit_classify").All(&direction)
	c.Data["classify"] = &direction
	c.TplName = "admin/tjkc.html"
}

type TjkcDataController struct {
	beego.Controller
}

//向数据库添加课程
//流程：1.如果接收不到图片信息则程序会返回。
//      2. 判断接收数据是否为空，如果有一个为空则返回JSON提示。
//      3. 保存图片到文件夹中，文件夹以当天日期为名，如果没有则创建。
//      4. 保存数据到数据库中。
func (c *TjkcDataController) Post() {
	f, h, err := c.GetFile("img")
	if err != nil {
		log.Fatal("getfile err ", err)
		return
	}
	//得到数据库模型
	ccit_course := models.Ccit_course{}
	//得到网页传输过来的数据
	ccit_course.Classify_id = c.GetString("Classify") //方向id
	ccit_course.Course_info = c.GetString("Synopsis") //课程介绍
	ccit_course.Course_name = c.GetString("Title")    //课程标题
	//判断数据是否符合要求
	if ccit_course.Course_name == "" || ccit_course.Course_info == "" || ccit_course.Classify_id == "" {
		c.Data["json"] = map[string]interface{}{"code": 2, "message": "请填写完整信息"}
		c.ServeJSON()
		return
	}
	//将上传的文件保存在指定的文件夹中
	c.SaveToFile("img", utils.IsDir("static/upload/"+time.Now().Format("2006-01-02"))+"/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	defer f.Close()                                                                                   //关闭流
	//向数据库保存数据
	o := orm.NewOrm()
	ccit_course.Courset_id = utils.Random()                                                              //课程id
	ccit_course.Putaway_time = time.Now().Format("2006-01-02")                                           //创建时间
	ccit_course.Course_path = "static/upvideo/" + ccit_course.Classify_id + "/" + ccit_course.Courset_id //课程的视频存放的文件夹
	ccit_course.Cover_img = "static/upload/" + time.Now().Format("2006-01-02") + "/" + h.Filename        // 拼接图片地址
	ccit_course.Is_permit = "0"                                                                          //是否上架，1 表示上架 、0表示不上架
	_, err = o.Insert(&ccit_course)
	if err == nil {
		//创建一个课程的视频存放的文件夹
		os.Mkdir("static/upvideo/"+ccit_course.Classify_id+"/"+ccit_course.Courset_id, os.ModePerm)
		//然后向创建的文件夹内放个文件，主要是描述一下这个文件夹是放什么的
		f, err1 := os.Create("static/upvideo/" + ccit_course.Classify_id + "/" + ccit_course.Courset_id + "/README.md")
		defer f.Close()
		if err1 != nil {
			//错误日志
		} else {
			f.Write([]byte("这个文件夹存放的是 " + ccit_course.Course_name + " 课程"))
		}
		//当前专业的课程数量 +1
		sql := "update ccit_classify set num=num+1 where classify_id = " + ccit_course.Classify_id
		classifly := []models.Ccit_classify{}
		//将上传的图片生成成缩略图
		err := utils.Thumbnail("static/upload/"+time.Now().Format("2006-01-02")+"/"+h.Filename, 216, 120)
		if err != nil {
			fmt.Println(err)
		}
		//执行
		o.Raw(sql).QueryRows(&classifly)
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "/admin/course"}
		c.ServeJSON()
		return
	} else {
		println(err)
	}
}

type TjzjPageController struct {
	beego.Controller
}

func (c *TjzjPageController) Get() {
	o := orm.NewOrm()
	direction := []models.Ccit_classify{}
	o.QueryTable("ccit_classify").All(&direction)
	c.Data["classify"] = &direction
	c.TplName = "admin/tjzj.html"
}

type TjUserController struct {
	beego.Controller
}

//接收文件 ->读取信息 ->循环解析数据并向数据库插入数据
func (c *TjUserController) Post() {
	//接收上传的excel表
	f, h, err := c.GetFile("uploadfile")
	//判断接收的文件是否有错误，如果有错误那么就记录在日志中并向前端反馈信息
	if err != nil {
		log.Fatal("getfile err ", err)
		c.Data["json"] = map[string]interface{}{"code": 2, "message": err}
		c.ServeJSON()
		return
	} else {
		//将上传的文件保存在指定的文件夹中。 保存位置在 static/upload, 没有文件夹要先创建  PS：这里没有做在保存文件时出现错误的处理
		c.SaveToFile("uploadfile", utils.IsDir("static/upload/excel")+"/"+h.Filename)
		defer f.Close() //关闭流
	}
	// 读取Excel 如果出现错误则退出。并删除文件
	if xlsx, err := excelize.OpenFile("static/upload/excel" + "/" + h.Filename); err != nil {
		fmt.Println(err)
		//如果打开这个文件出现了错误那么就删除这个文件。并向前端反馈信息
		os.Remove("static/upload/excel" + "/" + h.Filename)
		os.Exit(1)
		return
	} else {
		o := orm.NewOrm()
		//打开Excel 中的 Sheet1表
		rows, err := xlsx.GetRows("Sheet1")
		//获得一个结构体模型
		user := models.Ccit_user{}
		//遍历这张表格有多少行数据，并一行行循环取出
		for _, row := range rows {
			// 将对应的数据放进结构体中
			user.Student_number = row[0]              //这一行数据中的第一列数据
			user.Student_password = utils.MD5(row[1]) //这一行数据中的第二列数据 md5 加密
			user.Student_name = row[2]                //这一行数据中的第三列数据
			user.Student_phone = row[3]               //这一行数据中的第四列数据
			user.Student_activity = row[4]            //这一行数据中的第五列数据
			//判断格式是否正确  如果首行的格式不是 {学号 密码 姓名 手机号 是否激活（1激活/0不激活）} 这样的则格式错误
			//如果学号和手机号的长度是11位，那么写进数据库
			if len(user.Student_number) == 11 && len(user.Student_phone) == 11 {
				if _, err = o.Insert(&user); err != nil {
					//记录下向数据库插入信息出现错误的信息。 ps：待写
					return
				}
			}
		}
		//删除上传的文件
		//	os.Remove("static/upload/excel" + "/" + h.Filename)
		//通知前端将数据导入数据成功
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "/admin/course"} //ps：待定修改内容
		c.ServeJSON()
	}
}
