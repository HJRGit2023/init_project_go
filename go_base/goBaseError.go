package main

import (
	"fmt"
	"errors"
)

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
	return 0,nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
 }
// 自定义错误类型
 type DivideError struct {
	// 自定义错误类型 DivideError 包含两个字段，Dividend 和 Divisor，分别表示被除数和除数。
	Dividend int
	Divisor	int
 }
// 实现 error 接口类型
 func (d *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", d.Dividend, d.Divisor)
 }

 func divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{a, b}
	}
	return a / b, nil
}

// 通过实现 error 接口类型来生成错误信息。
func main() {
	err := errors.New("This is an error")
	fmt.Println(err)

	result, err := Sqrt(-10)
	if err != nil {
		fmt.Println(result,err)
	}

	result2, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result2)
	}

	_, err = divide2(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}