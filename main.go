package main

import (
	"log"
	"server/dao"
	"server/router"
	util "server/utils/config"
)

func main() {
	//读取配置文件 连接数据库
	m, err := util.LoadYamlConfig()
	if err != nil {
		log.Println("配置文件读取出错:", err)
		return
	}
	log.Println("yaml config loads successfully")

	//初始化redis连接池
	dao.InitRedis(m["redis"]["address"].(string), m["redis"]["password"].(string))
	log.Println("redis pool connect successfully")

	//开启路由
	router.SetupRouter()
	log.Println("server starts at port 8080")
}
