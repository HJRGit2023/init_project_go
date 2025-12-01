package main

import "fmt" 
/* 声明全局变量 */
var a int = 20

func main() {
	/* main 函数中声明局部变量 */
   var a int = 10
   var b int = 20
   var c int = 0
	fmt.Printf("main() 函数中 a = %d\n", a)
	c = sum(a, b)
	fmt.Printf("main() 函数中 c = %d\n", c)
}

/* 函数定义中的参数为  形参 */
func sum(a, b int) int {
	/* sum 函数中声明局部变量 */
	fmt.Printf("sum() 函数中 a = %d\n",  a)
	fmt.Printf("sum() 函数中 b = %d\n",  b)
	return a + b
}