package main

import "fmt"

const (
	// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
	// 所以 a=0, b=1, c=2 可以简写为 如下
	a = iota // 0
	b // 1	
	c // 2
	d = "ha" // 独立值，iota+=1
	e //  "ha" iota+=1
	f = 100 // iota += 1
	g // 100 iota += 1
	h = iota // 7,恢复计数
	i
	j=1<<iota // 相当于 1左移9位，100000000000，即 512
	k=3<<iota // 相当于 3左移10位，110000000000，即 3072
	l // 相当于 3左移11位，1100000000000，即6144 
	m // 相当于 3左移12位，11000000000000，即 12288
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m)
}