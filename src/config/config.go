package config

import "fmt"

// ConfigBean 配置文件结构体
type ConfigBean struct {
	inFilePath  string
	outFilePath string
}

// LoadConfig 载入配置文件
func LoadConfig() (cb ConfigBean) {
	fmt.Println("config loading...")

	var conf = ConfigBean{}

	conf.inFilePath = ""
	conf.outFilePath = ""

	fmt.Println("config end")

	return conf
}
