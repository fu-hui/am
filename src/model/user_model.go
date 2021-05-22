package model

type User struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}
