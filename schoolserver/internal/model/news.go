package model

type News struct {
	Id         int    `json:"Id" xml:"Id" form:"Id" query:"Id"`
	Title      string `json:"Title" xml:"Title" form:"Title" query:"Title"`
	Text       string `json:"Text" xml:"Text" form:"Text" query:"Text"`
	Preview    string `json:"Preview" xml:"Preview" form:"Preview" query:"Preview"`
	Created_at string `json:"Created_at" xml:"Created_at" form:"Created_at" query:"Created_at"`
}
