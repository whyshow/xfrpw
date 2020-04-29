package models

import (
	"github.com/astaxie/beego/orm"
)

//章节表映射
type Ccit_section struct {
	Section_id   string `orm:"column(section_id);pk"`
	Section_name string
	Courset_id   string
	Section_info string
	Section_part string
	Putaway_time string
}

//根据章节ID删除章节
func DeleteSection(ID string) bool {
	o := orm.NewOrm()
	if _, err := o.Delete(&Ccit_section{Section_id: ID}); err == nil {
		return true
	}
	return false
}
