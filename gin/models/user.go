package models

type User struct {
	Id       int    `form:"id"`
	Username string `form:"username"`
	Age      int    `form:"age"`
	Email    string `form:"email"`
	AddTime  int64  `form:"-"`
}

func (User) TableName() string {
	return "user"
}
