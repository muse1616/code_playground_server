package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
)

func Register(ctx *gin.Context) {
	//	获取邮箱 用户名 密码 验证码
	var r model.Register
	if err := ctx.Bind(&r); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg":        "服务器错误",
			"meta":       "401",
			"error_code": err.Error(),
		})
		return
	}
	result, msg, err := model.RegisterDao(r)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"result":     false,
			"msg":        msg,
			"meta":       http.StatusForbidden,
			"error_code": err.Error(),
		})
		return
	}
	if result == true {
		ctx.JSON(http.StatusOK, gin.H{
			"result":     true,
			"msg":        msg,
			"meta":       http.StatusOK,
			"error_code": "",
		})
		return
	} else if result == false {
		ctx.JSON(http.StatusOK, gin.H{
			"result":     false,
			"msg":        msg,
			"meta":       http.StatusOK,
			"error_code": "",
		})
		return
	}
}
