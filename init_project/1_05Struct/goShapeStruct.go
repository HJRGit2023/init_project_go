package main

import "fmt"
/* 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。 */

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	r float64
}

func (r Rectangle) Area() {
	fmt.Printf("Rectangle area: %f\n", r.width*r.height)
}

func (r Rectangle) Perimeter() {
	fmt.Printf("Rectangle perimeter: %f\n", 2*(r.width+r.height))
}

func (c Circle) Area() {
	fmt.Printf("Circle area: %f\n", 3.14*c.r*c.r)
}

func (c Circle) Perimeter() {
	fmt.Printf("Circle perimeter: %f\n", 2*3.14*c.r)
}

func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{r: 3}
	r.Area()
	r.Perimeter()
	c.Area()
	c.Perimeter()
}