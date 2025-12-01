package main

import "fmt"

func main() {
	// -------------------算术运算符-----------------------
	// 自增与自减只能以 <var name>++ 或者 <var name>-- 的模式声明，
	// 并且只能单独存在，不能  在  自增或自减   的同时做  加减乘除  的计算
	a := 1
	// 正确写法
	a++
	a--

	// 错误的使用方式
	// ++a
	//--a

	// 错误使用方式，不可以自增时计算,也不能赋值
	//b := a++ + 1
	//c := a--


	d := 10 + 0.1
    e := byte(1) + 1
    fmt.Println(d, e)
	// 不同类型不能直接相加，需要转换类型
    sum := d + float64(e)
    fmt.Println(sum)

    sub := byte(d) - e
    fmt.Println(sub)

    mul := d * float64(e)
    div := int(d) / int(e)

    fmt.Println(mul, div)
	// -------------------关系运算符-------------------
	// 关系运算符结果只会是 bool 类型。
	aa := 1
	bb := 5

	fmt.Println("aa == bb: ",aa == bb)
	fmt.Println("aa != bb: ",aa != bb)
	fmt.Println("aa  > bb: ", aa > bb)
	fmt.Println("aa  < bb: ",aa < bb)
	fmt.Println("aa >= bb: ", aa >= bb)
	fmt.Println("aa <= bb: ", aa <= bb)
	// -------------------逻辑运算符------------
	aaa := true
	bbb := false

	fmt.Println("aaa && bbb: ", aaa && bbb)
	fmt.Println("aaa || bbb: ", aaa || bbb)
	fmt.Println("!(aaa && bbb): ", !(aaa && bbb))
	// -------------------位运算符----------------
	fmt.Println("0位与0位: ", 0 & 0)
	fmt.Println("0位或0位: ", 0 | 0)
	fmt.Println("0位异或0位: ", 0 ^ 0)

	fmt.Println("0位与1位: ", 0 & 1)
	fmt.Println("0位或1位: ",0 | 1)
	fmt.Println("0位异或1位: ", 0 ^ 1)

	fmt.Println("1位与1位: ", 1 & 1)
	fmt.Println("1位或1位: ", 1 | 1)
	fmt.Println("1位异或1位: ", 1 ^ 1)

	fmt.Println("1位与0位: ", 1 & 0)
	fmt.Println("1位或0位: ", 1 | 0)
	fmt.Println("1位异或0位: ", 1 ^ 0)
}