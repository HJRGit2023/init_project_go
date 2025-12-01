package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)
	go func() {
		for i :=0;i<100;i++ {
			// fmt.Println("发送数据：", i)
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println("接收数据：", v)
	}
	time.Sleep(time.Second*5)

}