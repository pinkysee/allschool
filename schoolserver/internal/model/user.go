package model

type User struct {
	ID        int64  `json:"ID" xml:"ID" form:"ID" query:"ID"`
	Name      string `json:"Name" xml:"Name" form:"Name" query:"Name"`
	Login     string `json:"Login" xml:"Login" form:"Login" query:"Login"`
	Password  string `json:"Password" xml:"Password" form:"Password" query:"Password"`
	Role      string `json:"Role" xml:"Role" form:"Role" query:"Role"`
	Classname string `json:"Classname" xml:"Classname" form:"Classname" query:"Classname"`
	Avatar    int    `json:"Avatar" xml:"Avatar" form:"Avatar" query:"Avatar"`
}
