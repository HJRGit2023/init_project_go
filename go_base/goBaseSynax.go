package main

import "fmt"

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	vname1 int
	vname2 int
)

func main(){
	// 第一种：声明变量并初始化

	var a string = "Runoob"
	fmt.Println(a)

	var b,c int = 10,20
	fmt.Println("int 类型的变量 b,c = ",b,c)
	// 声明一个变量但不初始化 int类型默认值是0
	var d int
	fmt.Println("int 类型的变量 d = ", d)

	// 声明一个变量但不初始化 bool类型默认值是false
	var e bool
	fmt.Println("bool 类型的变量 e = ",e)

	// 第二种：根据类型自动判断变量的类型
	var f = true
	fmt.Println("bool 类型的变量 f = ",f)

	// 第三种：省略类型声明，根据变量的值自动判断变量的类型
	// v_name := value

	// 这种也是一种声明，:=也是声明
	// intVal := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	// intVal := 1 相等于：下面两句
	// var intVal int 
	// intVal =1


    var intVal int
	// intVal := 100 // 这时候上一句+这句 会产生编译错误，因为 intVal 已经声明，不需要重新声明
	
	fmt.Println("根据变量的值自动判断变量的类型 intVal = ", intVal)

	// 可以将 var f string = "Runoob" 简写为 f := "Runoob"：
	g := "Runoob"
	fmt.Println("string 类型的变量 f = ",g)


	// 声明多个变量, 多次声明会编译错误
	// var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断
//vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误

	// 类型相同多个变量, 非全局变量
	var vname1, vname2, vname3 int = 1, 2, 3
	fmt.Println("类型相同多个变量, 非全局变量 vname1, vname2, vname3 = ", vname1, vname2, vname3)
}