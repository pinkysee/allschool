package model

type Test struct {
	Id          int    `json:"Id" xml:"Id" form:"Id" query:"Id"`
	Name        string `json:"Name" xml:"Name" form:"Name" query:"Name"`
	Quetion     string `json:"Quetion" xml:"Quetion" form:"Quetion" query:"Quetion"`
	Teachername string `json:"Teachername" xml:"Teachername" form:"Teachername" query:"Teachername"`
	Lesson      string `json:"Lesson" xml:"Lesson" form:"Lesson" query:"Lesson"`
	Komu        string `json:"Komu" xml:"Komu" form:"Komu" query:"Komu"`
	Deadline    string `json:"Deadline" xml:"Deadline" form:"Deadline" query:"Deadline"`
	Answer      string `json:"Answer" xml:"Answer" form:"Answer" query:"Answer"`
	Rightans    string `json:"Rightans" xml:"Rightans" form:"Rightans" query:"Rightans"`
}
