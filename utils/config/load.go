package util

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

//读取yaml配置文件 本项目直接读取为map 不使用unmarshal
func LoadYamlConfig() (result map[string]map[string]interface{}, err error) {
	result = make(map[string]map[string]interface{})
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &result)
	if err != nil {
		return
	}
	return
}
