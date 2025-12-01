package main

import (
	"fmt"
	"time"
)
/* 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数 */
func printOdd() {
	for i := 1; i <= 10; i++ {
		if i%2!= 0 {
			fmt.Println("奇数：", i)
		}
	}
}

func printEven() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("偶数：", i)
		}
	}
}

func main() {
	go printOdd()
	go printEven()
	// 等待一段时间确保所有goroutine完成,   否则主线程会立即退出
    time.Sleep(time.Second)
}