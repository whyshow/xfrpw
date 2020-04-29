package models

import (
	"github.com/astaxie/beego/orm"
)

//小节表映射
type Ccit_knobble struct {
	Knobble_id    string `orm:"column(knobble_id);pk"`
	Section_id    string
	Courset_id    string
	Knobble_name  string
	Video_address string
	Section_part  string
}

/**
 * 根据小节ID查询小节信息
 */
func QueryKnobble(knobble_id string) (bool, string, Ccit_knobble) {
	o := orm.NewOrm()
	knobble := Ccit_knobble{Knobble_id: knobble_id}
	err := o.Read(&knobble)
	if err == orm.ErrNoRows {
		return false, "查询不到", knobble
	} else if err == orm.ErrMissPK {
		return false, "找不到主键", knobble
	} else {
		return true, "查询成功", knobble
	}
}

/**
 * 根据章节ID和章节号查询当前章节的小节信息
 * 返回数据行数，小节信息切片
 */
func QueryKnobbles(section_id string, section_part string) (int64, []Ccit_knobble) {
	o := orm.NewOrm()
	knobble := []Ccit_knobble{}
	num, err := o.Raw("SELECT * FROM ccit_knobble WHERE Section_id = ? AND Section_part = ?", section_id, section_part).QueryRows(&knobble)
	if err == nil {
		return num, knobble
	} else {
		return 0, nil
	}
}

/**
 * 根据小节ID删除单个小节
 */
func DeleteKnobble(knobble_id string) (bool, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Ccit_knobble{Knobble_id: knobble_id}); err == nil {
		if num > 0 {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, err
	}
}

/**
 *  根据章节ID和章节号来删除当前章节号的所有小节
 */
func DeleteKnobbles(section_id string, section_part string) (bool, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Ccit_knobble{Section_id: section_id, Section_part: section_part}); err == nil {
		if num > 0 {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, err
	}
}
