package main

import "fmt"

func main() {
	var s string = "Go语言"
	var bytes []byte = []byte(s)
	var runes []rune = []rune(s)
	// 直接获取字符串的长度 len(s) 是把 字符串转换成 []byte 之后 , 再获取 []byte 的长度
	fmt.Println("string length: ", len(s))
	fmt.Println("bytes length: ", len(bytes))
	// 字符串 转换成 []rune，[]rune 的长度  则是  字符串字符  的长度
	fmt.Println("runes length: ", len(runes))

	fmt.Println("string sub: ", s[0:7])
    fmt.Println("bytes sub: ", string(bytes[0:7]))
    fmt.Println("runes sub: ", string(runes[0:3]))
}