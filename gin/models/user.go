package models

type User struct {
	Id       int    `form:"id" json:"id"`
	Username string `form:"username" json:"username"`
	Age      int    `form:"age" json:"age"`
	Email    string `form:"email" json:"email"`
	AddTime  int64  `form:"-" json:"addTime"`
}

func (User) TableName() string {
	return "user"
}

type UserArgument struct {
	Username string `form:"username"`
}
