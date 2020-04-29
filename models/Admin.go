package models

import (
	"../utils"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"net/http"
	"time"
)

var log_info *logs.BeeLogger

func init() {
	log_info = utils.Init_info_log()
}

//用户 表映射
type Ccit_admin struct {
	Admin_id         int `orm:"column(admin_id);pk"`
	Admin_account    string
	Admin_password   string
	Register_time    time.Time
	Admin_permission string
}

func AdminLogin(ccitUser Ccit_admin, password string, request *http.Request) int64 {
	o := orm.NewOrm()
	err := o.Read(&ccitUser, "Admin_account")
	if err != nil {
		log_info.Informational("信息：" + "用户 " + ccitUser.Admin_account + " 登录失败" + " 原因：" + "用户名不存在" + " " + "登录IP地址：" + utils.ClientIP(request))
		log_info.Flush()
		return 1
	}
	if ccitUser.Admin_password != password {
		log_info.Informational("信息：" + "用户 " + ccitUser.Admin_account + " 登录失败" + " 原因：" + "密码错误" + " " + "登录IP地址：" + utils.ClientIP(request))
		log_info.Flush()
		return 2
	} else {
		log_info.Informational("信息：" + "用户 " + ccitUser.Admin_account + " 登录成功" + " " + "登录IP地址：" + utils.ClientIP(request))
		log_info.Flush()
		return 0
	}
}

func AdminiExit(username string, request *http.Request) {
	log_info.Informational("信息：" + "用户 " + username + " 退出登录" + " 原因：" + "管理员手动退出" + " " + "退出IP地址：" + utils.ClientIP(request))
	log_info.Flush()
}
