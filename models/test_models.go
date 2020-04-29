package models

//章节表映射
type Ccit_section0 struct {
	Section_id   string `orm:"column(section_id);pk"`
	Section_name string
	Courset_id   string
	Section_info string
	Section_part string
	Putaway_time string
	Ccit_knobble []Ccit_knobble0
}

//小节表映射
type Ccit_knobble0 struct {
	Knobble_id   string
	Section_id   string
	Courset_id   string
	Knobble_name string
	//Video_address string
	Section_part string
	//Ccit_section *Ccit_section `orm:"rel(fk)"`
}
