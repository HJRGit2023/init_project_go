package main

import "fmt"

func main() {
	var a int = 20 /* 声明实际变量 */
	var ip *int /* 指向整型 声明指针变量 */
	if(ip != nil) {
		fmt.Printf("ip 的值为%x ，ip 不是空指针\n", ip)
	}
	if(ip == nil) {
		fmt.Printf("ip 的值为%x ，ip 是空指针\n", ip)
	}
	/* Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址 */
	ip = &a
	var fp *float32 /* 指向浮点型 */
	var sp *string /* 指向字符串 */
	fmt.Printf("a变量的地址为：%p\n", &a)
	/* 指针变量的存储地址 */
    fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	fmt.Printf("fp 变量储存的指针地址: %x\n", fp )
	if(fp != nil) {
		fmt.Printf("fp 的值为%x ，fp 不是空指针\n", fp)
	}
	if(fp == nil) {
		fmt.Printf("fp 的值为%x ，fp 是空指针\n", fp)
	}
	fmt.Printf("sp 变量储存的指针地址: %x\n", sp )
	/* 指针变量的实际值 */
	fmt.Printf("ip 变量指向的实际值：%d\n", *ip)

}