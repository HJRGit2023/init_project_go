package main

import "fmt"
/* 题目：判断一个整数是否是回文数 */
func isPalindromic(num int) bool {
	if num <0 {
		fmt.Println(num, "is not a palindromic number")
		return false
	}
	numStr := fmt.Sprintf("%d", num)
	for i := 0; i < len(numStr)/2; i++ {
		// 首尾字符不相等
		if numStr[i] != numStr[len(numStr)-i-1] { 
			fmt.Println(num, "is not a palindromic number")
			return false
		}
	}
	fmt.Println(num, "is a palindromic number")
	return true
}

func main() {
	num := 12321
	num1 := 123321
	num2 := 12344421
	num3 := -1
	num4 := 0
	isPalindromic(num)
	isPalindromic(num1)
	isPalindromic(num2)
	isPalindromic(num3)
	isPalindromic(num4)
}