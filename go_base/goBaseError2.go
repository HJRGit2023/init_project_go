package main

import "fmt"

// 定义一个 `DivideError` 结构体，用于记录除数为零的错误信息
type DivideError struct {
	Dividend int
	Divisor int
}
// 实现 `Error()` 方法，用于返回错误信息
func (e *DivideError) Error() string {
	strFormat := `Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0`
	return fmt.Sprintf(strFormat, e.Dividend)
}
// 定义 `int` 类型除法运算的函数
func divide(varDividend int, varDivisor int) (result int, errMessage string) {
	if varDivisor == 0 {
		data := DivideError{Dividend:varDividend, Divisor:varDivisor,}
		errMessage = data.Error()
		return
	} else {
		return varDividend / varDivisor, ""
	}
}

func main() {
	// 正常情况
	if result, errMessage := divide(100, 5); errMessage == "" {
		fmt.Println("Result:", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errMessage := divide(100, 0); errMessage != "" {
		fmt.Println("error Message is :", errMessage)
	}
}