package main

import "fmt"
/* 1 composition 组合继承 
	组合是 Go 中实现代码复用的主要方式。通过将一个结构体嵌入到另一个结构体中，
	子结构体可以"继承"父结构体的字段和方法。
2 interface inheritance 接口继承 
3 完整继承模拟*/
// 1.1 父结构体 2.1 接口继承
type Animal struct {
	Name string
}
// 1.2 父结构体的方法 2.2 接口方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "says meow")
}

// 1.3 子结构体 2.3 接口实现
type Dog struct {
	Animal // 嵌入 Animal 结构体
	Breed string
}

// 2.1 接口继承
type Speaker interface {
	Speak()
}
// 3.1 完整继承模拟 交通工具，车辆；
type Vehicle struct {
	Brand string
}
// 3.2 完整继承模拟 交通工具的方法
func (v *Vehicle) Start() {
	fmt.Println(v.Brand, "started")
}
// 3.3 完整继承模拟 车辆结构体
type Car struct {
	Vehicle // 嵌入 Vehicle 结构体
	Model string
}
// 3.4 完整继承模拟 车辆的方法 重写父结构体的方法Start
func (c *Car) Start() {
	fmt.Println(c.Brand, c.Model,"car started")
}

func main() {
	// 1.4 创建 Dog 结构体
	dog := Dog{Animal: Animal{"Buddy"}, Breed: "Golden Retriever"} //金毛猎犬
	dog.Speak() // 1.5 调用父结构体的方法
	fmt.Println(dog.Name, "Breed is a", dog.Breed)

	// 2.1 接口继承
	var speaker Speaker = &dog // 2.2 接口实现
	speaker.Speak() // 2.3 调用接口方法

	// 3.1 完整继承模拟
	v := Vehicle{Brand: "BMW"}
	c := Car{Vehicle: Vehicle{Brand: "Toyota"}, Model: "Civic"}

	v.Start() // 3.5 调用父结构体的方法
	c.Start() // 3.6 调用重写父结构体的方法
	c.Vehicle.Start() // 3.7 调用嵌入的父结构体的方法


}
