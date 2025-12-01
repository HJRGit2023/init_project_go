package main
 
import (
	"fmt"
	"time"
	"sync"
)
/* 1 go 启动 Goroutine，调用sayHello函数，并传入参数"World" 
 2 通道（Channel）是用于 Goroutine 之间的数据传递。
 3 定义一个可以存储整数类型的带缓冲通道
 4 go 遍历通道和关闭通道
 5 select 语句使得一个 goroutine 可以等待多个通信操作。select 会阻塞，直到其中的某个 case 可以继续执行
 6 sync.WaitGroup 用于等待多个 Goroutine 完成。*/

// 1.1 sayHello 函数定义 
func sayHello(name string) {
	for i := 0; i< 5;i++ {
		fmt.Println("Hello", name)
		time.Sleep(100*time.Millisecond)
	}
}

// 2.1 定义一个函数，用于计算数组的和，并将结果通过通道（Channel）返回
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将结果通过通道（Channel）返回
}

// 4.1 斐波那契数列 传入通道，关闭通道
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 关闭通道
}

// 5.1 select 语句使得一个 goroutine 可以等待多个通信操作。
// select 会阻塞，直到其中的某个 case 可以继续执行。
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// 6.1 sync.WaitGroup 用于等待多个 Goroutine 完成。
func worker(id int,wg *sync.WaitGroup) {
	// defer 表示 “延迟执行”，即函数执行完毕后（无论正常结束还是 panic），
	// 会调用 wg.Done()，作用是将 WaitGroup 的计数器减 1，标记当前协程已完成。
	defer wg.Done() // 等待计数器减1 Goroutine 完成时调用 Done()
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	// 1.2go 启动 Goroutine，调用sayHello函数，并传入参数"World"
	go sayHello("World")
	// 主线程休眠100毫秒，等待Goroutine执行完毕
	// 你会看到输出的 Main 和 Hello。输出是没有固定先后顺序，因为它们是两个 goroutine 在执行：
	for i:=0;i < 5;i++{
		fmt.Println("Main")
		time.Sleep(100* time.Millisecond)
	}

	// 2.2 定义一个数组，并将其传递给 sum 函数，并接收返回的结果
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 接收两个结果
	fmt.Println(x, y, x+y)

	// 3.1 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch :=make(chan int,2)
	// 3.2 向通道中写入数据
	ch <- 1
	ch <- 2
	// 3.3读取通道中的数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 4.2 斐波那契数列
	c = make(chan int, 10)
	go fibonacci(10, c)
	// 4.2.1 range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 4.2.2 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 4.2.3 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 4.2.4 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}

	// 5.2 斐波那契数列 传入通道，关闭通道
	c1 := make(chan int)
	quit := make(chan int)
	// 5.2.1 向通道中写入数据
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c1)
		}
		quit <- 0 // 关闭通道
	}()
	// fibonacci2 goroutine 在 channel c1 上发送斐波那契数列，当接收到 quit channel 的信号时退出。
	fibonacci2(c1, quit) // 接收数据

	// 6.2 多个 Goroutine 等待
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) // 计数器加1
		go worker(i, &wg) // 启动 Goroutine
	}
	wg.Wait() // 等待计数器减1 Goroutine 完成时调用 Done()
	fmt.Println("All workers finished")
}