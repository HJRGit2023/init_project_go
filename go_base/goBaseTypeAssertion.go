package main

import (
	"fmt"
)
/* 类型断言（Type Assertion）是一种用于检查接口值的实际类型的机制 
value, ok := interfaceValue.(Type)
	其中：
    interfaceValue 是一个接口类型的变量。
    Type 是你想要断言的类型。
    value 是断言成功后的具体类型的值。
    ok 是一个布尔值，表示断言是否成功。
*/

func printType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("这是一个整数:", v)
    case string:
        fmt.Println("这是一个字符串:", v)
    default:
        fmt.Println("未知类型")
    }
}

func processInterface(i interface{}) {
    if s, ok := i.(string); ok {
        fmt.Println("处理字符串:", s)
    } else if n, ok := i.(int); ok {
        fmt.Println("处理整数:", n)
    } else {
        fmt.Println("无法处理的类型")
    }
}

func main() {
	var i interface{} = "Hello World"
	s, ok := i.(string)
	// 尝试将 i 断言为 string 类型
	if ok {
		fmt.Println("断言成功:", s)
	} else {
		fmt.Println("断言失败")
	}
	// 尝试将 i 断言为 int 类型
	a, ok := i.(int)
	if ok {
		fmt.Println("断言成功:", a)
	} else {
		fmt.Println("断言失败")
	}

	// 直接断言为 string 类型
	s1 := i.(string)
	fmt.Println("直接断言成功:", s1)
	// 直接断言 为 int 类型（会引发 panic）
	// a1 := i.(int)
	// fmt.Println("直接断言成功:", a1)

	// switch 语句使用类型断言
	printType(100)
	printType("Hello World")
	printType(true)

	processInterface(100)
	processInterface("Hello World")
	processInterface(true)
}

