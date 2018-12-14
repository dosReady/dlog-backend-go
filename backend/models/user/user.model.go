package user

import (
	. "github.com/dosReady/dlog/backend/models"
	dao "github.com/dosReady/dlog/backend/modules/dao"
)

type DlogUser struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserCall     string `json:"user_call"`
	Base
}

func UserList() *[]DlogUser {
	users := make([]DlogUser, 0)
	conn := dao.GetConnection()
	conn.Find(&users)
	conn.Close()
	return &users
}
