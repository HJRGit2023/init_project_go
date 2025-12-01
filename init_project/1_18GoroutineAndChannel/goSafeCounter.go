package main

import (
	"fmt"
	"sync"
	"time"
)
/* 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值 */
type SafeCounter struct {
	mu sync.Mutex
	count int
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	sc.count++
	sc.mu.Unlock()
}

func (sc *SafeCounter) GetVal() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

func main() {
	sc := SafeCounter{}
	// 启动10个goroutine
	for i := 0; i<10; i++ {
		go func() {
			// 每个goroutine执行1000次Increment
			for j := 0; j<1000; j++ {
				sc.Increment()
			}
		}()
	}

	// 等待一段时间确保所有goroutine完成
    time.Sleep(time.Second)

    // 输出最终计数
    fmt.Printf("Final count: %d\n", sc.GetVal())

}

