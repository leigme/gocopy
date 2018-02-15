package main

import (
	"config"
	"fmt"
)

func main() {
	fmt.Println("app start...")

	config.LoadConfig()

	fmt.Println("app end!")
}
