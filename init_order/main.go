package main

import (
	"fmt"
	_ "github.com/learn/init_order/pkg1"
)

const mainName = "main"

var mainVar = getMainName()

func init() {
    fmt.Println("main init method invoked")
}

func getMainName() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}

func main() {
	fmt.Println("main method invoked!")
}

