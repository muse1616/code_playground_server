package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/session"
	"server/utils"
)

//权限认证中间件
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		//取cookie
		id, err1 := c.Cookie("session_id")
		str, err2 := c.Cookie("xxx")
		//根据id取值
		userId, err3 := session.GetSession(id)
		if err1 != nil || err2 != nil || err3 != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"msg":   "登录状态失效,请重新登录",
				"meta":  http.StatusForbidden,
				"error": "",
			})
			c.Abort()
			return
		}
		//判断sessionId是否有效
		decode, err := utils.Decode(str)
		fmt.Println("str:", decode)
		fmt.Println("userid:", userId)
		if err == nil && decode == userId {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"msg":   "登录状态失效,请重新登录",
				"meta":  http.StatusForbidden,
				"error": "",
			})
			c.Abort()
			return
		}

	}
}
