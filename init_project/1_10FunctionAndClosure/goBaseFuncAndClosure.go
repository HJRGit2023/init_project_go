package	main

import "fmt"

type A struct {
    i int
}
// 定义方法 是一个包含 接受者 的函数，大部分情况下可以通过类型的实例调用。
func (a *A) add(v int) int {
    a.i += v
    return a.i
}

// 声明函数变量
var function1 func(int) int

// 声明闭包
var squart2 func(int) int = func(p int) int {
    p *= p
    return p
}

func main() {
    a := A{1}
    // 把方法a.add  赋值  给   函数变量function1
    function1 = a.add
    
    // 声明一个闭包并直接执行
    // 此闭包返回值是另外一个闭包（带参闭包）
    returnFunc := func() func(int, string) (int, string) {
		// 此时参数是func(int, string) (int, string)
        fmt.Println("this is a anonymous function")
		// anonymous 匿名的，不知名的；不记名的，不具名的；无特色的
        return func(i int, s string) (int, string) {
            return i, s
        }
    }()

    // 执行returnFunc闭包并传递参数
    ret1, ret2 := returnFunc(1, "test")
    fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)

    fmt.Println("a.i = ", a.i)
    fmt.Println("after call function1, a.i = ", function1(1))
    fmt.Println("a.i = ", a.i)
}