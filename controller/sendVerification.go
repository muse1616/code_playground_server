package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	ConfigUtil "server/utils/config"
	EmailUtil "server/utils/email"
)

func SendVerification(ctx *gin.Context) {

	//接受参数
	email := ctx.PostForm("email")

	//读取配置文件
	configMap, err := ConfigUtil.LoadYamlConfig()
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg":   "服务器错误",
			"meta":  "400",
			"error": err.Error(),
		})
		return
	}

	//发送验证码
	code, err := EmailUtil.SendEmail(email, configMap["email"]["username"].(string), configMap["email"]["password"].(string), configMap["email"]["host"].(string))

	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg":   "验证码发送失败",
			"meta":  "401",
			"error": err.Error(),
		})
		return
	}

	//	redis缓存验证码
	if err := EmailUtil.SaveEmailVerification(email, code); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg":   "验证码发送失败",
			"meta":  "401",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "验证码发送成功",
		"meta":  "400",
		"error": "",
	})

}
