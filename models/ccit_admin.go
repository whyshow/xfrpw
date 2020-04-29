package models

import (
	_ "github.com/go-sql-driver/mysql"
)

//专业方向表映射
type Ccit_classify struct {
	Classify_id   string `orm:"column(classify_id);pk"`
	Classify_name string
	Add_date      string
	Indexes       string
	Num           int
	Classify_path string
}

//这个是课程表映射，但是它加入了 专业方向表里面的 Classify_name，Indexes字段。
type Ccit_course0 struct {
	Courset_id    string `orm:"pk"`                  //课程id
	Classify_id   string `orm:"column(classify_id)"` //专业方向id
	Course_name   string //课程名称
	Course_info   string //课程信息
	Cover_img     string //课程封面地址
	Putaway_time  string //上架时间
	Label         string //标签
	Is_permit     string //是否上架
	Classify_name string
	Indexes       string
}

//课程表映射
type Ccit_course struct {
	Courset_id   string `orm:"column(courset_id);pk"` //课程id
	Classify_id  string //专业方向id
	Course_name  string //课程名称
	Course_info  string //课程信息
	Cover_img    string //课程封面地址
	Putaway_time string //上架时间
	Label        string //标签
	Is_permit    string //是否上架
	Course_path  string //课程路径

}

//课程表映射
type Ccit_courseAndCcit_classify struct {
	Course_name   string `orm:"column(course_name);pk"` //课程名称
	Course_info   string //课程信息
	Classify_name string
	Indexes       string
}

// Iflysse@2018
// aiNI@13141
func init() {
	//mac 下打包linux CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	//mac 下打包windows CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
	// set GOARCH=amd64
	// set GOOS=linux
	// go build
	// chmod +x    赋予执行权限
	// nohup ./XFRPW
}
