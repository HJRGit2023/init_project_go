package main

import "fmt"
/* 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。 */

func add(n *int) {
	*n += 10 // 计算  指针指向的值  并增加10
}

func main() {
	var num int
	num = 10
	fmt.Println("Before modification:", num)
	add(&num) // 调用函数，传入指针,修改指针指向的值
	fmt.Println("After modification:", num)
}