package main

import "fmt"

func main() {
	/*以下实例声明一个名为 numbers 的整数数组，其大小为 5，
	在声明时，数组中的每个元素都会根据其数据类型进行默认初始化，
	对于整数类型，初始值为 0。*/
    var arr [5]int
	for i:=0; i<len(arr); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
	/* 还可以使用初始化列表来初始化数组的元素： 
	声明一个大小为 5 的整数数组，并将其中的元素分别初始化为 1、2、3、4 和 5。*/
	var number1 [5]int = [5]int{1, 2, 3, 4, 5}
	for i:=0; i<len(number1); i++ {
		fmt.Printf("number1[%d] = %d\n", i, number1[i])
	}
	/* 还可以使用 := 简短声明语法来声明和初始化数组 */
	number2 := [5]int{1, 2, 3, 4, 5}
	for i:=0; i<len(number2); i++ {
		fmt.Printf("number2[%d] = %d\n", i, number2[i])
	}
	/* 如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度： */
	var number3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出数组元素 */ 
	for i:=0; i<len(number3); i++ {
		fmt.Printf("number3[%d] = %f\n", i, number3[i])
	}

	/* 如果设置了数组的长度，我们还可以通过指定下标来初始化元素： */
	balance := [5]float32{1:2.0, 3:7.0}
	/* 输出数组元素 */ 
	for i:=0; i<len(balance); i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	/* 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小： */
	balance[4] = 50.0

	/* 数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值 */
	var salary float32 = balance[4]
	fmt.Printf("salary = %f\n", salary)

	var n [10]int /* n 是一个长度为 10 的数组 */
	var j,h int
	/* 为数组 n 初始化元素 */        
	for j=0; j<10; j++ {
		n[j] = j + 100 /* 设置元素为 j + 100 */
	}
	/* 输出每个数组元素的值 */
	for h=0; h<10; h++ {
		fmt.Printf("Element[%d]: %d\n", h, n[h])
	}



}