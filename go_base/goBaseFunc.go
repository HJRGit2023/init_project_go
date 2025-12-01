package main

import "fmt"

func main() {
	fmt.Println(max(10, 20))

	/* 定义局部变量 */
   var a int = 100
   var b int = 200
   var ret int

   /* 调用函数并返回最大值 */
   ret = max1(a, b)

   fmt.Printf( "最大值是 : %d\n", ret )
}

func max(num1, num2 int) int {
	/* 声明局部变量 */

	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

/* 函数返回两个数的最大值 */
func max1(num1, num2 int) int {
   /* 定义局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}