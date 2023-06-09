package core

import (
	"fmt"
	"gin/config"
	"gin/global"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func InitConf() {
	//加载logo
	loadLogo()
	//配置全局变量
	global.Config = readConf()
	//log输出
	log.Printf("%v: yaml配置文件加载成功!", global.Config.Prefix)
}

// 读取yaml配置文件
func readConf() *config.Config {
	//读取命令行配置文件地址
	option := *global.Option
	//配置文件地址
	var yamlPath = "settings.yaml"
	if option.Setting != "" {
		yamlPath = option.Setting
	}

	//初始配置文件体
	config := &config.Config{}

	//加载配置文件流
	yamlConfig, readErr := ioutil.ReadFile(yamlPath)
	if readErr != nil {
		//将错误信息格式化字符串输出
		panic(fmt.Errorf("读取yaml配置文件失败:%s", readErr))
	}

	//解析yaml文件
	yamlErr := yaml.Unmarshal(yamlConfig, config)
	if yamlErr != nil {
		log.Fatalf("解析yaml配置文件失败:%v", yamlErr)
	}

	return config
}

// 加载logo
func loadLogo() {
	logoFilePath := "banner.txt"

	//读取logo文件
	logo, readErr := ioutil.ReadFile(logoFilePath)
	if readErr != nil {
		//将错误信息格式化字符串输出
		return
	}
	fmt.Println(string(logo))
}
