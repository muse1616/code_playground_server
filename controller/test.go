package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type TModel struct {
	user_name string `json:"user_name"`
}

func Test(ctx *gin.Context) {
	var t TModel
	if err := ctx.Bind(&t); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"result": false,
			"msg":    "服务器错误",
			"meta":   http.StatusForbidden,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "登录状态有效",
		"meta":  "400",
		"error": "null",
	})
	return
}
