package main

import (
	"fmt"
	"math"
)

// 定义一个简单接口 Shape
type Shape interface {
	/* Shape 是一个接口，定义了两个方法：Area 和 Perimeter。
	任意类型只要实现了这两个方法，就被认为实现了 Shape 接口*/
    Area() float64
    Perimeter() float64
}

// 定义一个结构体
type Circle struct {
    Radius float64
}

// 实现 Shape 接口的 Area 方法, 计算圆的面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现 Shape 接口的 Perimeter 方法, 计算圆的周长
func (c Circle) Perimeter() float64 {
        return 2 * math.Pi * c.Radius
}

func printValue(value interface{}) {
	// 空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集
	fmt.Printf("value type is %T, value is %v\n", value, value)
}

func main() {
    // 声明一个 Circle 类型的变量
	var c Circle = Circle{Radius: 5}
	// 接口变量可以存储实现了接口的类型, c 实现了 Shape 接口，因此可以赋值给接口变量 s
	var s Shape = c 
	fmt.Println("The area of the circle is:", s.Area())
	fmt.Println("The perimeter of the circle is:", s.Perimeter())

	printValue(42) // int
	printValue("hello") // string
	printValue(c) // Circle 类型实现了 Shape 接口
	printValue(s) // Circle 类型实现了 Shape 接口
	printValue(3.1415926) // float64
	printValue([]int{1, 2, 3}) // 切片类型 slice
}

