package main

import "fmt"

func main() {
	// ----------不使用 append() 函数时，函数内部会对  切片进行修改
	s := make([]int, 3, 6)
    fmt.Println("s length:", len(s))
    fmt.Println("s capacity:", cap(s))
    fmt.Println("initial, s = ", s)
    s[1] = 2
    fmt.Println("set position 1, s = ", s)
	//在不使用 append() 函数的情况下，在 函数内部  对  切片的修改，
	// 都会影响   到    原始实例
    modifySlice(s)
    fmt.Println("after modifySlice, s = ", s)
	// ------------使用 append()函数时, 没有触发扩容操作，只是在  原有切片  的基础上，
	s1 := make([]int, 3, 6)
	s2 := append(s1, 4)
    fmt.Println("after append, s2 length:", len(s2))
    fmt.Println("after append, s2 capacity:", cap(s2))
    fmt.Println("after append, s1 =", s1)
    fmt.Println("after append, s2 =", s2)

    s1[0] = 1024
    fmt.Println("after set position 0, s1 =", s1)
    fmt.Println("after set position 0, s2 =", s2)
	// 在函数内部使用 append() 函数，会修改  原始实例，s1 和 s2 都会受到影响
    appendInFunc(s1)
    fmt.Println("after append in func, s1 =", s1)
    fmt.Println("after append in func, s2 =", s2)
	// 触发扩容操作的 append() 函数
	appendInFunc2(s1)
    fmt.Println("after append in func2扩容, s1 =", s1)
    fmt.Println("after append in func2扩容, s2 =", s2)
}

func modifySlice(param []int) {
    param[0] = 1024
}
// 在函数内部使用 append() 函数，会修改  原始实例，s1 和 s2 都会受到影响
func appendInFunc(param []int) {
	// param 长度为 3，容量为 6，未改变容量
    param = append(param, 1022)
    fmt.Println("in func, param =", param)
    param[2] = 512
    fmt.Println("set position 2 in func, param =", param)
}
// 触发扩容操作的 append() 函数
func appendInFunc2(param []int) {
	// 入参s1长度为3，容量为6，s1作为param传入，
	// 触发扩容操作，改变了param的容量，并不会改变s1的值
    param = append(param, 511)
	param = append(param, 512)
	param = append(param, 513)
	param = append(param, 514)
    param2 := append(param, 515)
    fmt.Println("in func, param1 =", param)
    param2[2] = 500
    fmt.Println("set position 2 in func, param2 =", param2)
}