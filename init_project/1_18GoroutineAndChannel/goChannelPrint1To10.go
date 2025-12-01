package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("发送数据：", i, "到通道 ", time.Now().Format("2006-01-02 15:04:05.000"))
		ch <- i
	}
	close(ch)
}

func recvData(ch chan int) {
	timeOut := time.After(time.Second * 10)
	for {
		select {
			case data, ok := <-ch:
				if !ok {
					fmt.Println("通道关闭")
					return
				}
				fmt.Println("接收数据：", data, "从通道 ", time.Now().Format("2006-01-02 15:04:05.000"))
			case <-timeOut:
				fmt.Println("超时")
			default:
				fmt.Println("无数据，等待中。。。")
				time.Sleep(time.Second*2)
		}
	}
}

func main() {
	ch := make(chan int, 10)
	go sendData(ch)
	go recvData(ch)
	time.Sleep(time.Second * 10)
}