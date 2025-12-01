package main

import "fmt"

func main() {
		/* ------------------------赋值运算符------------------------ */
	// 赋值运算符用于给变量赋值，可以将右侧的值赋给左侧的变量。
	// 赋值运算符包括：=, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=。
	// 赋值运算符的优先级低于算术运算符，因此在进行赋值运算时，先进行算术运算，再进行赋值运算。

    a, b := 1, 2
    var c int
    c = a + b
    fmt.Println("c = a + b, c =", c)

    plusAssignment(c, a)
    subAssignment(c, a)
    mulAssignment(c, a)
    divAssignment(c, a)
    modAssignment(c, a)
    leftMoveAssignment(c, a)
    rightMoveAssignment(c, a)
    andAssignment(c, a)
    orAssignment(c, a)
    norAssignment(c, a)
	/* ------------------------其他运算符-------------- */
	aaa := 4
    var ptr *int
    fmt.Println(aaa)

    ptr = &aaa
    fmt.Printf("*ptr 为 %d\n", *ptr) // 输出 *ptr 的值
	/* ----------------------运算优先级------------------------- */
	// 运算符的优先级决定了运算顺序，运算符的优先级高的运算先进行。
	// 运算符的优先级从高到低依次为：
	// 1. 圆括号 ()	// 2. 乘除法 **	// 3. 加减法 + -	// 4. 位移 << >>
	// 5. 关系运算符 < <= > >=	// 6. 相等运算符 ==!=	// 7. 位运算符 & ^ |
	// 8. 逻辑运算符 && || !	// 9. 赋值运算符 = += -= *= /= %= <<= >>= &= ^= |=
	// 10. 逗号运算符 ,	// 11. 函数调用	// 12. 索引 []	// 13. 切片 [:]
	// 14. 结构体字段 .	// 15. 指针 *	// 16. 通道操作 <-	// 17. 类型转换
	// 18. 其他运算符 ...
	var aa int = 21
    var bb int = 10
    var cc int = 16
    var d int = 5
    var e int

    e = (aa + bb) * cc / d // ( 31 * 16 ) / 5
    fmt.Printf("(aa + bb) * cc / d 的值为 : %d\n", e)

    e = ((aa + bb) * cc) / d // ( 31 * 16 ) / 5
    fmt.Printf("((aa + bb) * cc) / d 的值为  : %d\n", e)

    e = (aa + bb) * (cc / d) // 31 * (16/5)
    fmt.Printf("(aa + bb) * (cc / d) 的值为  : %d\n", e)

    // 21 + (160/5)
    e = aa + (bb*cc)/d
    fmt.Printf("aa + (bb * cc) / d 的值为  : %d\n", e)

    // 2 & 2 = 2; 2 * 3 = 6; 6 << 1 = 12; 3 + 4 = 7; 7 ^ 3 = 4;4 | 12 = 12
    f := 3 + 4 ^ 3 | 2&2*3<<1
    fmt.Println(f == 12)

}

func plusAssignment(c, a int) {
    c += a // c = c + a
    fmt.Println("c += a, c =", c)
}

func subAssignment(c, a int) {
    c -= a // c = c - a
    fmt.Println("c -= a, c =", c)
}

func mulAssignment(c, a int) {
    c *= a // c = c * a
    fmt.Println("c *= a, c =", c)
}

func divAssignment(c, a int) {
    c /= a // c = c / a
    fmt.Println("c /= a, c =", c)
}

func modAssignment(c, a int) {
    c %= a // c = c % a
    fmt.Println("c %= a, c =", c)
}

func leftMoveAssignment(c, a int) {
    c <<= a // c = c << a
    fmt.Println("c <<= a, c =", c)
}

func rightMoveAssignment(c, a int) {
    c >>= a // c = c >> a
    fmt.Println("c >>= a, c =", c)
}

func andAssignment(c, a int) {
    c &= a // c = c & a
    fmt.Println("c &= a, c =", c)
}

func orAssignment(c, a int) {
    c |= a // c = c | a
    fmt.Println("c |= a, c =", c)
}

func norAssignment(c, a int) {
    c ^= a // c = c ^ a
    fmt.Println("c ^= a, c =", c)
}