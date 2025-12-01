package main

import (
	"fmt"
	"time"
)

// channel作为参数，限制channel操作的方向，仅发送数据 <-在chan的右边，表示只能发送数据
func sendOnly(ch chan <- int) {
	for i := 0; i < 5; i++ {
        ch <- i
        fmt.Printf("发送: %d\n", i)
    }
    close(ch)
}

// channel作为参数，限制channel操作的方向，仅接收数据 <-在 chan 左边，表示只能接收数据
func receiveOnly(ch <- chan int) {
	for v := range ch {
        fmt.Printf("接收到: %d\n", v)
    }
}

func main() {
	// 创建一个带缓冲的channel，容量为3
	ch := make(chan int, 3)
	// 启动发送goroutine
    go sendOnly(ch)

    // 启动接收goroutine
    go receiveOnly(ch)
	// 使用select进行多路复用
	timeOut := time.After(time.Second * 2)
	for {
		select {
			case v, ok :=<-ch: 
				if !ok {
					fmt.Println("Channel已关闭")
					return
				}
				fmt.Printf("主goroutine接收到: %d\n", v)
			case <-timeOut:
				fmt.Println("超时退出")
				return
			default:
				fmt.Println("没有数据，等待中...")
            	time.Sleep(500 * time.Millisecond)
		}
	}

}