package main

import (
	"config"
	"fmt"
)

func main() {
	fmt.Println("app start...")

	cb := config.LoadConfig()

	fmt.Println(cb.InFilePath)
	fmt.Println(cb.OutFilePath)

	fmt.Println("app end!")
}
