package model

import "reflect"

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}
