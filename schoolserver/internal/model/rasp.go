package model

type Rasp struct {
	Classname string `json:"classname" xml:"classname" form:"classname" query:"classname"`
	Lessons   string `json:"lessons" xml:"lessons" form:"lessons" query:"lessons"`
}
