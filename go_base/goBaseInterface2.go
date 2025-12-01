package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (n NokiaPhone) call() {
	fmt.Println("Calling from NokiaPhone")
}

type IPhone struct {
}

func (i IPhone) call() {
	fmt.Println("Calling from IPhone")
}

type Shape interface {
	area() float64
}

type Rectangle struct { // 矩形
	width float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct { // 圆形
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()

	var shape Shape
	shape = Rectangle{width :10, height : 20}
	fmt.Printf("矩形面积Rectangle area:%f\n", shape.area())
	shape = Circle{10}
	fmt.Printf("圆形面积Circle area:%f\n", shape.area())
}
