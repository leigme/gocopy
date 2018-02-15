package main

import (
	"config"
	"fmt"
)

func main() {
	fmt.Println("app start...")

	cb := config.LoadConfig()
	copyFiles(cb.InFilePath, cb.OutFilePath)

	fmt.Println("app end!")
}
