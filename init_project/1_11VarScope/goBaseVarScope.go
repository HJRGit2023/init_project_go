package main

import (
	"fmt"
	"time"
)
var aa int = 1
func main() {
	// -------------------------局部变量的作用域------------
	var a int
	// if 语句里的变量声明，只在if语句块内有效
    if b := 1; b == 0 {
        fmt.Println("b == 0")
    } else {
        c := 2
        fmt.Println("declare c = ", c)
        fmt.Println("b == 1")
    }
    // fmt.Println(b)
    // fmt.Println(c)
	// switch语句里的变量声明，只在case语句块内有效
    switch d := 3; d {
    case 1:
        e := 4
        fmt.Println("declare e = ", e)
        fmt.Println("d == 1")
    case 3:
        f := 4
        fmt.Println("declare f = ", f)
        fmt.Println("d == 3")
    }
    // fmt.Println(e)
	// for语句里的变量声明，只在for语句块内有效
    for i := 0; i < 1; i++ {
        forA := 1
        fmt.Println("forA = ", forA)
    }
    // fmt.Println("forA = ", forA)
	// select语句里的变量声明，只在case语句块内有效
    select {
    case <-time.After(time.Second):
        selectA := 1
        fmt.Println("selectA = ", selectA)
    }
    // fmt.Println("selectA = ", selectA)

    // 匿名代码块
    {
        blockA := 1
        fmt.Println("blockA = ", blockA)
    }
    // fmt.Println("blockA = ", blockA)

    fmt.Println("a = ", a)
	// -------------------------全局变量的作用域------------
	{
        fmt.Println("global variable, aa = ", aa)
        aa = 3 // 修改全局变量的值
        fmt.Println("global variable, aa = ", aa)

        aa := 10 // 声明局部变量aa，覆盖全局变量aa
        fmt.Println("local variable, aa = ", aa)
        aa-- // 局部变量aa的值--
        fmt.Println("local variable, aa = ", aa)
    }
    fmt.Println("global variable, aa = ", aa)
	// 优先  使用作用域更小  的变量的规则，同样适用于 局部变量 
	var bb int = 4
    fmt.Println("local variable, bb = ", bb)
	// if 语句里bb的声明，只在if语句块内有效，
	// if里 声明了bb:=3,if 会执行，bb的值为3，然后执行if语句块，bb的值--，然后执行else语句块，bb的值为2
    if bb := 3; bb == 3 {
		// 
        fmt.Println("if statement, bb = ", bb)
        bb--
        fmt.Println("if statement, bb = ", bb)
    }
	// 局部变量 bb= 4，if里的bb变为2，出了if失效了，所以bb的值还是4
    fmt.Println("local variable, bb = ", bb)
}