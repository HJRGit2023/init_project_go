package main

import (
	"fmt"
	"time"
)
/* 设计一个任务调度器，接收一组任务（可以用函数表示），
并使用协程并发执行这些任务，同时统计每个任务的执行时间 */

func scheduler(tasks []func()) {
	for i, task := range tasks {
		go func(i int, task func()) {
			fmt.Println("Task", i, "finished", time.Now().Format("2006-01-02 15:04:05.000"))
		}(i, task)
	}
}

func main() {
	// tasks := []func(){}
	// 启动100个goroutine同时增加计数
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println("Task", i, "started", time.Now().Format("2006-01-02 15:04:05.000"))
		}()
	}
	// scheduler(tasks)
	time.Sleep(4 * time.Second)
}