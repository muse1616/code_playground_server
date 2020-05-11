package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	ConfigUtil "server/utils/config"
)

var (
	DB *gorm.DB
)

//初始化mysql
func InitMysql() (err error) {
	//读取配置文件
	configMap, err := ConfigUtil.LoadYamlConfig()
	if err != nil {
		return
	}
	//dsn
	dsn := configMap["mysql"]["username"].(string) + ":" + configMap["mysql"]["password"].(string) + "@tcp(" + configMap["mysql"]["address"].(string) + ")/" + configMap["mysql"]["dbName"].(string) + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(dsn)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//测试连通
	err = DB.DB().Ping()
	if err != nil {
		return err
	}
	return
}

// 关闭mysql连接
func CloseMysqlConnection() (err error) {
	err = DB.Close()
	return
}
