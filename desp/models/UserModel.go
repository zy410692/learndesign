package models

type UserCreateFunc func(id int, username string) interface{}

type User struct {
	Id       int
	UserName string
}

func NewUser() UserCreateFunc {
	return func(id int, username string) interface{} {
		return &User{Id: id, UserName: username}
	}
}

type Admin struct {
	Id        int
	AdminName string
	Role      string
}

func NewAdmin() UserCreateFunc {
	return func(id int, username string) interface{} {
		return &Admin{Id: id, AdminName: username, Role: "Admin"}
	}
}
