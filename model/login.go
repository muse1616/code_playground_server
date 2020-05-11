package model

import (
	"server/dao"
)

// 登录模型 account为账号或者邮箱
type LoginUser struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func LoginDao(user LoginUser) (result bool, username string, email string) {
	//处理登录请求
	dao.DB.SingularTable(true)
	dao.DB.AutoMigrate(&User{})
	//先判断是否为用户名登录
	var u User
	dao.DB.Where("username = ? AND password = ?", user.Account, user.Password).First(&u)
	if u.Password == user.Password && u.Username == user.Account {
		result = true
		username = u.Username
		email = u.Email
		return
	}
	//再判断是否为邮箱登录
	dao.DB.Where("email = ? AND password = ?", user.Account, user.Password).First(&u)
	if u.Password == user.Password && u.Email == user.Account {
		result = true
		username = u.Username
		email = u.Email
		return
	}

	result = false
	return
}
