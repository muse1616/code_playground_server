package router

import (
	"github.com/gin-gonic/gin"
	"server/controller"
	"server/middleware"
)

func SetupRouter() *gin.Engine {
	//默认路由
	r := gin.Default()
	//v1路由组
	v1Group := r.Group("v1")
	{
		//注册验证码发送
		v1Group.POST("/register/verification", controller.SendVerification)
		//注册路由
		v1Group.POST("/register/confirm", controller.Register)
		//登录路由 用户名或邮箱
		v1Group.POST("/login", controller.Login)
		//测试路由 测试登录权限
		v1Group.POST("/auth/test", middleware.LoginAuth(), controller.Test)
	}
	//开启路由
	err := r.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	//返回路由实例
	return r
}
