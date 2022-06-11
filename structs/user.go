package structs

type User struct {
	Name    string `form:"user" json:"user" xml:"user" binding:"required"`
	Address string `form:"address" json:"address" xml:"address" binding:"required"`
	Age     int    `form:"age" json:"age" xml:"age" binging:"required"'`
}
