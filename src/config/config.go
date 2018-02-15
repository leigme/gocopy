package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// ConfigBean 配置文件结构体
type ConfigBean struct {
	InFilePath  string
	OutFilePath string
}

// LoadConfig 载入配置文件
func LoadConfig() (cb *ConfigBean) {
	fmt.Println("config loading...")
	file, err := os.Open("config.json")
	if err != nil {

		panic(err)
	}
	defer file.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("exception: file failed...")
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}

	configBean := &ConfigBean{}

	err = json.Unmarshal(GetValidByte(chunks), &configBean)

	if err != nil {
		fmt.Println("exception: json failed...")
		panic(err)
	}

	fmt.Println("config end")
	return configBean
}

// GetValidByte 去除空字符串
func GetValidByte(src []byte) []byte {
	var str_buf []byte
	for _, v := range src {
		if v != 0 {
			str_buf = append(str_buf, v)
		}
	}
	return str_buf
}
