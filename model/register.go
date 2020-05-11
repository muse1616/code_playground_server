package model

import (
	"fmt"
	"server/dao"
	"server/utils/email"
)

//注册表单模型
type Register struct {
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	Password     string `json:"password"`
	Verification string `json:"verification"`
}

//数据库user表
type User struct {
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

//注册处理函数
func RegisterDao(r Register) (result bool, msg string, err error) {
	//先判断验证码 再判断email 最后判断username
	//根据email读取验证码
	code, err := email.ReadVerificationByEmail(r.Email)
	if err != nil || code != r.Verification {
		result = false
		msg = "验证码错误"
		return
	}
	//检测邮箱是否已注册

	dao.DB.SingularTable(true)
	dao.DB.AutoMigrate(&User{})
	var u User
	dao.DB.Where("email = ?", r.Email).First(&u)
	if u.Email == r.Email {
		result = false
		msg = "该邮箱已被注册"
		return
	}
	dao.DB.Where("username = ?", r.UserName).First(&u)
	if u.Username == r.UserName {
		result = false
		msg = "该用户名已被注册"
		fmt.Println("该用户名已被注册")
		return
	}
	//注册
	u = User{
		Username: r.UserName,
		Password: r.Password,
		Email:    r.Email,
	}
	if err = dao.DB.Create(&u).Error; err != nil {
		result = false
		msg = "注册失败"
		return
	}
	result = true
	msg = "注册成功"
	return
}
