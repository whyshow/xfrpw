package controllers

import (
	"../models"
	"../utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"strings"
	"time"
)

type UpZipController struct {
	beego.Controller
}

//接收并处理上传的压缩包文件  接收->保存 ->解压 ->删除压缩包 ->循环改名并向数据库写数据->通知前台信息
func (c *UpZipController) Post() {
	//接收web前端上传的视频压缩包
	f, h, err := c.GetFile("uploadZip")
	section := models.Ccit_section{}
	knobble := models.Ccit_knobble{}
	Classify := c.GetString("Classify")
	section.Courset_id = c.GetString("course")
	section.Section_part = c.GetString("part")
	section.Section_name = c.GetString("Title")
	section.Section_info = c.GetString("Synopsis")
	section.Putaway_time = time.Now().Format("2006-01-02") //创建时间
	// 判断数据是否为空   -> ps:为空则发生提示信息给前端  Ajax方式
	if section.Courset_id == "" || Classify == "" || section.Section_part == "" || section.Section_name == "" || section.Section_info == "" || err != nil {
		fmt.Println("getfile err ", err)
		c.TplName = "admin/tjzj.html"
		//上传的信息不完整
		return
	} else {
		section.Section_id = utils.Random()
		o := orm.NewOrm()
		section1 := []models.Ccit_section{}
		//SQL语句
		sql := "SELECT * FROM ccit_section WHERE courset_id =" + section.Courset_id + " AND section_part =" + c.GetString("part")
		if r, _ := o.Raw(sql).QueryRows(&section1); r != 0 {
			// 打印日志：这个章节已经存在了
			c.TplName = "admin/tjzj.html"
			return
		} else {
			//保存接收上传的压缩包
			err1 := c.SaveToFile("uploadZip", utils.IsDir("static/upvideo/"+Classify+"/"+section.Courset_id+"/"+section.Section_part+"/")+h.Filename) // 保存位置在 static/upvideo, 没有文件夹要先创建
			defer f.Close()                                                                                                                           //关闭流
			if err1 != nil {
				fmt.Println("保存错误 ", err1)
				c.TplName = "admin/tjzj.html"
				return
			} else {
				//记录开始解压前的时间
				start := time.Now()
				//解压  参数 一：要解压的文件所处于的位置及名称。参数二：解压后存放的位置
				if err2 := utils.DeCompress("static/upvideo/"+Classify+"/"+section.Courset_id+"/"+section.Section_part+"/"+h.Filename, "static/upvideo/"+Classify+"/"+section.Courset_id+"/"+section.Section_part+"/"); err2 != nil {
					fmt.Println("解压错误错误 ", err2)
					c.TplName = "admin/tjzj.html"
					return
				} else {
					//打印解压所属要的时间
					fmt.Println("整个解压需要的时间 ", time.Since(start))
					//删除压缩包
					if os.Remove("static/upvideo/"+Classify+"/"+section.Courset_id+"/"+section.Section_part+"/"+h.Filename) != nil {
						//日志 删除错误

						//  从这里开始往下一旦出现错误就删除这个整个章节的所有视频，包括文件夹
						c.TplName = "admin/tjzj.html"
						return
					} else {
						//向数据库写入章节信息
						if _, err := o.Insert(&section); err != nil {
							// 打印错误日志
							c.TplName = "admin/tjzj.html"
							return
						}
						namearr := []string{}
						namearr, _ = utils.GetAllFile("static/upvideo/"+Classify+"/"+section.Courset_id+"/"+section.Section_part, namearr)
						for num := range namearr {
							//生成视频的id
							video := utils.Random()
							//修改视频的名字
							err := os.Rename("static/upvideo/"+Classify+"/"+section.Courset_id+"/"+section.Section_part+"/"+namearr[num], "static/upvideo/"+Classify+"/"+section.Courset_id+"/"+section.Section_part+"/"+video+string(namearr[num][strings.LastIndex(namearr[num], "."):len(namearr[num])]))
							if err != nil {
								//如果重命名文件失败,则输出错误 file rename Error!
								fmt.Println("file rename Error!")
								//打印错误详细信息
								fmt.Printf("%s", err)
								c.TplName = "admin/tjzj.html"
							} else {
								//循环向数据库写入小节信息
								knobble.Knobble_id = video                                                          //视频id
								knobble.Section_id = section.Section_id                                             //章节id
								knobble.Knobble_name = string(namearr[num][0:strings.LastIndex(namearr[num], ".")]) //小节的名字
								knobble.Courset_id = section.Courset_id                                             //课程id
								knobble.Section_part = section.Section_part                                         //章节
								knobble.Video_address = "static/upvideo/" + Classify + "/" + section.Courset_id + "/" + section.Section_part + "/" + video + string(namearr[num][strings.LastIndex(namearr[num], "."):len(namearr[num])])
								if _, err := o.Insert(&knobble); err != nil {
									// 打印错误日志
									fmt.Println(err)
									c.TplName = "admin/tjzj.html"
								}
							}
						}
						c.TplName = "admin/section.html"
						return
					}
				}
			}
		}
	}
}
