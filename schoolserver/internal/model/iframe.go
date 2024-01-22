package model

type Iframe struct {
	ID   int64  `json:"ID" xml:"ID" form:"ID" query:"ID"`
	Page string `json:"Page" xml:"Page" form:"Page" query:"Page"`
	Path string `json:"Path" xml:"Path" form:"Path" query:"Path"`
}
