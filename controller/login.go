package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/session"
	"server/utils"
)

func Login(ctx *gin.Context) {
	var loginUser model.LoginUser
	if err := ctx.Bind(&loginUser); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"result": false,
			"msg":    "服务器错误",
			"meta":   http.StatusForbidden,
			"error":  err.Error(),
		})
		return
	}
	result, username, email := model.LoginDao(loginUser)
	if result == false {
		ctx.JSON(http.StatusForbidden, gin.H{
			"result": false,
			"msg":    "账号或密码错误",
			"meta":   http.StatusOK,
			"error":  "",
		})
		return
	}
	//登录成功

	//处理session_id
	sessionId := utils.GenerateUUid()
	_ = session.SetSession(sessionId, username)
	ctx.SetCookie("session_id", sessionId, 3600*24, "/", "127.0.0.1", false, true)
	//加密用户名
	encoded := utils.Encode(username)
	ctx.SetCookie("xxx", encoded, 3600*24, "/", "127.0.0.1", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"result":   true,
		"msg":      "登录成功",
		"username": username,
		"email":    email,
		"meta":     http.StatusOK,
		"error":    "",
	})
	return

}
