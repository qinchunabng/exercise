package entity

//定义Member结构体
type Member struct {
	Name string `form:"name" json:"name" binding:"required,NameValid"`
	Age  int    `form:"age" json:"age" binding:"required,gte=10,lt=120" msg:"年龄必须大于10小于120"`
}
