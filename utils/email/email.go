package util

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"server/dao"
	"strings"
	"time"
)

/**
邮箱发送验证码
*/
func SendEmail(email string, username string, password string, host string) (code int, err error) {

	auth := smtp.PlainAuth("", username, password, host)
	to := []string{email}
	user := username
	nickname := "Code Playground"
	subject := "【Code Playground】邮箱验证"
	contentType := "Content-Type: text/plain; charset=UTF-8"

	//生成验证码
	rand.Seed(time.Now().Unix())
	code = rand.Intn(9000) + 1000

	body := fmt.Sprintf("验证码:%d,有效时间:30分钟\r\n.", code)
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err = smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		return
	}
	return
}

/**
保存验证码至redis
*/
func SaveEmailVerification(email string, code int) (err error) {
	conn := dao.Pool.Get()
	//ping
	err = dao.Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return
	}
	//设置
	_, err = conn.Do("Set", email, code)
	if err != nil {
		return err
	}
	//30分钟过期
	_, err = conn.Do("expire", email, 1800)
	if err != nil {
		return err
	}
	return
}
