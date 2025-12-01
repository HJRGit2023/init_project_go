package main

import (
	"fmt"
	"time"
	"sync"
)
// 安全计数器
type SafeCounter struct {
	mu sync.Mutex
	count int
}

// 不安全计数器
type UnsafeCounter struct {
	count int
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	sc.count++
	sc.mu.Unlock()
}

func (sc *SafeCounter) GetCount() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

func (uc *UnsafeCounter) Increment() {
	uc.count++
}

func (uc *UnsafeCounter) GetCount() int {
	return uc.count
}



func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("routine say : ", 500*i, "，毫秒后说：", s, " ", time.Now())
	}
}

func main() {
	go func() {
		// closure 闭包，在闭包中运行线程
		fmt.Println("run goroutine in closure")
	}()
	go func(string) {
	}("gorouine: closure params")
	go say("in goroutine: world")
	say("in main: hello")

	// counter := UnsafeCounter{}
	counter := SafeCounter{}

    // 启动100个goroutine同时增加计数
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待一段时间确保所有goroutine完成
    time.Sleep(time.Second)

    // 输出最终计数
    fmt.Printf("Final count: %d\n", counter.GetCount())
}
