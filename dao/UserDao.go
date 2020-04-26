package dao

import (
	"github.com/lqlkxk/gin/initDB"
	"github.com/lqlkxk/gin/model"
)

type UserDao struct {
}

func (u UserDao) FindByname(name string) (user model.User) {
	queryUserStr := "select * from user where username=?"
	initDB.Db.Get(&user, queryUserStr, name)
	return user
}
