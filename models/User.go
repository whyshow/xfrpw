package models

import "github.com/astaxie/beego/orm"

type Ccit_user struct {
	Student_number   string `orm:"column(student_number);pk"`
	Student_password string
	Student_name     string
	Student_phone    string
	Student_activity string
}

/**
 * 用户登录
 */
func Login(user Ccit_user) (bool, int64, string, Ccit_user) {
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM ccit_user WHERE student_number = ? AND student_password = ? ", user.Student_number, user.Student_password).QueryRow(&user)
	if err == orm.ErrNoRows {
		if err := o.Raw("SELECT * FROM ccit_user WHERE student_number = ? ", user.Student_number).QueryRow(&user); err == orm.ErrNoRows {
			return false, -1, "账号不存在", Ccit_user{}
		} else {
			return false, 0, "密码错误", Ccit_user{}
		}
	} else {
		return true, 1, "登录成功", user
	}
}

/**
 * 找回用户密码
 */
func FindUserPassword(user Ccit_user) (bool, string) {
	o := orm.NewOrm()
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		return false, "账号或手机号错误"
	} else if err == orm.ErrMissPK {
		return false, "找不到主键"
	} else {
		//更新新的密码
		if _, err := o.Update(&user, "Student_password"); err == nil {
			return true, "修改成功"
		} else {
			return false, "数据库错误"
		}
	}
}

/**
 * 查询单个用户信息
 */
func QueryUserInfo(user Ccit_user) (bool, string, Ccit_user) {
	o := orm.NewOrm()
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		return false, "账号或手机号错误", user
	} else if err == orm.ErrMissPK {
		return false, "找不到主键", user
	} else {
		return true, "查询成功", user
	}
}

/**
 * 查询所有用户信息
 */
func QueryUserInfos() (bool, int64, []Ccit_user) {
	o := orm.NewOrm()
	user := []Ccit_user{}
	if num, err := o.Raw("SELECT * FROM ccit_user").QueryRows(&user); err == nil {
		return true, num, user
	} else {
		return false, num, user
	}
}
