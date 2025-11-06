package main

import "fmt"

type ConnState int
// iota 独立作用于  每个 const 定义组，就是上面看到的 const ( ``// ``code... ) 结构。
const (
	StateNew ConnState = iota // 0 iota 出现在const组中第一行，表示从0开始计数，即n-1
	StateActive
	StateIdle
	StateHijacked
	StateClosed
)

type Month int
const (
  January Month = 1 + iota
  February
  March
  April
  May
  June
  July
  August
  September
  October
  November
  December
)
// 一定要注意 iota 出现在定义组中的  第几行，而  不是  当前代码中它 第几次 出现，否则会导致计数错误。
const pre int = 1
const a int = iota
const (
    b int = iota
    c
    d
    e
)
const (
    f = 2
    g = iota
    h
    i
)

func main() {
	fmt.Println("ConnState StateNew = ", StateNew)
	fmt.Println("ConnState StateActive = ", StateActive)

	fmt.Println("一月 ：",January)
	fmt.Println("二月 ：", February)

	fmt.Println("a = ", a)
	fmt.Println("iota出现在const组中第一行，b = ", b)
	fmt.Println("iota出现在const组中第一行，c = ", c)

	fmt.Println("iota出现在const组中第二行，f = ", f)
	fmt.Println("iota出现在const组中第二行，g = ", g)

}