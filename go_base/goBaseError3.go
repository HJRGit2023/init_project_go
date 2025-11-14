package main

import (
	"fmt"
	"errors"
)
/* 1 errors.Is检查某个错误是否是特定错误或由该错误包装而成。
2 errors.As尝试将错误转换为特定类型。
3 Go 的 panic 用于处理不可恢复的错误，recover 用于从 panic 中恢复。
 */

var ErrNotFound = errors.New("not found")
// 1.1 判断错误是否是特定错误。
func FindIdItem(id int) error {
	return fmt.Errorf("database error: %w", ErrNotFound)
}

// 2.1 将错误转换为特定类型以便进一步处理。
type MyError struct {
	Code int
	Msg string
}

// 2.2 尝试将错误转换为MyError类型。
func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

// 2.3 定义一个函数返回MyError类型。
func GetError() error {
	return &MyError{
		Code: 404,
		Msg: "not found",
	}
}

// 3.1 定义一个函数，用于处理panic。
func safeFunction() {
	defer func() {
		if err := recover(); err!=nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()
	// 处理一些可能出现的panic。
	panic("Something went wrong")
}

func main() {
	err := FindIdItem(100)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Item not found")
	} else {
		fmt.Println("Other error:", err)
	}
	// 2.4 尝试将GetError()返回的错误转换为MyError类型。
	err2 := GetError()
	// 2.5 声明一个空指针变量myErr。指向MyError类型。
	var myErr *MyError
	// 2.6 使用errors.As尝试将err2转换为MyError类型。
	if errors.As(err2, &myErr) { // 2.7 如果成功转换，则myErr不为空。
	// &myErr表示将myErr的地址传给errors.As。
		fmt.Printf("Custome error-Code:%d,Msg:%s\n", myErr.Code, myErr.Msg)
	}

	// 3.2 调用safeFunction()函数，触发panic。
	fmt.Println("Starting program...")
	safeFunction()
	fmt.Println("Program continued after panic")

}