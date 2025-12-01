package main

import "fmt"
import "unsafe"

func main() {
	var p1 *int
	var p2 *string
	i := 10
	s := "Hello World"
	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
	// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	p1 = &i // 获取i变量的存储地址
	p2 = &s // 获取s变量的存储地址

	p3 := &p2 // 获取p2变量的存储地址

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	// 指针类型数据，可以通过指针 间接访问  其指向的变量
	fmt.Println(*p1 == i)
	*p1 = 2
	fmt.Println(i)

	a := 2
	var p *int // 声明一个指针变量p 
	fmt.Println(&a)
	p = &a // 指针变量p 指向 a变量  的存储地址
	fmt.Println(p, &a)

	var pp **int // 声明一个指向指针 的  指针变量
	pp = &p // 指针变量pp 指向 指针变量p 指向的 变量的 存储地址
	fmt.Println(pp, p)
	**pp = 3 // 间接修改 指针变量p 指向的 变量的值
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p) // *p表示间接访问 指针变量p 指向的 变量的值
	fmt.Println(a, &a)

	// 指针转换成 unsafe.Pointer:
	b := "Hello, world!"
	// uintptr 类型是把  内存地址  转换成了一个整数类型，可以用来做指针运算
	// unsafe.Pointer 类型是把  任意类型  转换成  指针类型，可以用来做  类型转换
	upA := uintptr(unsafe.Pointer(&b))
	// 指针运算：向后移动一个字节
	upA += 1
	// 类型转换：把 uintptr 类型转换成 指针类型
	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c, c)
}