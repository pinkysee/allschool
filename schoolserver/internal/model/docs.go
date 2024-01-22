package model

type Docs struct {
	Name string `json:"Name" xml:"Name" form:"Name" query:"Name"`
	Path string `json:"Path" xml:"Path" form:"Path" query:"Path"`
	Page string `json:"Page" xml:"Page" form:"Page" query:"Page"`
}
