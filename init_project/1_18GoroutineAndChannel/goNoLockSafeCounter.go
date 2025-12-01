package main

import (
	"fmt"
	"sync/atomic"
	"sync"
	// "time"
)

type Counter struct {
	count int64
}
// 
func (c *Counter) Increment() {
	// 原子操作加1
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) GetVal() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	counter := Counter{}
	// 并发执行10个goroutine，每个goroutine执行1000次counter.Increment()
	var wg sync.WaitGroup // 等待组
	for i := 0; i < 10; i++ {
		wg.Add(1) // 增加等待组计数
		go func() {
			defer wg.Done() // 减少等待组计数
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	// fmt.Println("Waiting for goroutines to finish...")
	// time.Sleep(time.Second)
	wg.Wait() // 等待所有goroutine执行完毕
	fmt.Println("最终计数：", counter.GetVal())

}