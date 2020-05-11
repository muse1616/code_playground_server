package main

import (
	"log"
	"server/dao"
	"server/router"
	util "server/utils/config"
)

func main() {
	/**
	配置文件
	*/
	m, err := util.LoadYamlConfig()
	if err != nil {
		log.Println("配置文件读取出错:", err)
		return
	}
	log.Println("yaml config loads successfully")

	/**
	数据库初始化
	*/
	//初始化redis连接池
	dao.InitRedis(m["redis"]["address"].(string), m["redis"]["password"].(string))
	log.Println("redis pool connect successfully")
	//初始化mysql数据库
	if err := dao.InitMysql(); err != nil {
		log.Println(err)
		return
	}
	log.Println("mysql connect successfully")

	/**
	开启路由
	*/
	router.SetupRouter()
	log.Println("server starts at port 8080")
}
