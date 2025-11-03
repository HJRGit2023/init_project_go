package main

import (
	"fmt"
	// "context"
	// "sync/atomic"
	// "time"
)

func main() {
	for i:=0; i<5; i++ {
		fmt.Println("方式一：i的下标是：", i)
	}
	// b := 1
	// for b < 10 {
	// 	// 无限循环
	// 	fmt.Println("方式二：b的值是：", b)
	// }
	// 方式三：无限循环
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	// var started bool
    // var stopped atomic.Bool
    // for {
    //     if !started {
    //         started = true
    //         go func() {
    //             for {
    //                 select {
    //                 case <-ctx.Done():
    //                     fmt.Println("ctx done")
    //                     stopped.Store(true)
    //                     return
    //                 }
    //             }
    //         }()
    //     }
    //     fmt.Println("main")
    //     if stopped.Load() {
    //         break
    //     }
    // }

	// 遍历数组
    var a [3]string
    a[0] = "Hello"
    for i := range a {
        fmt.Println("数组当前下标：", i)
    }
    for i, e := range a {
        fmt.Println("数组a[", i, "] = ", e)
    }

	// 遍历切片 容量为3
    s := make([]string, 3)
    s[0] = "Hello"
    for i := range s {
        fmt.Println("切片当前下标：", i)
    }
    for i, e := range s {
        fmt.Println("切片s[", i, "] = ", e)
    }
    
    
    m := make(map[string]string)
    m["b"] = "Hello, b"
    m["a"] = "Hello, a"
    m["c"] = "Hello, c"
    for i := range m {
        fmt.Println("map当前key：", i)
    }
    for k, v := range m {
        fmt.Println("m[", k, "] = ", v)
    }
}