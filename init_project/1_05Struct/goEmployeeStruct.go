package main

import "fmt"
/* 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息 */
type Person struct {
	Name string
	Age int
}

type Employee struct {
	EmployeeID int
	Person // 组合 Person 结构体，并添加 EmployeeID 字段，
}

func (e *Employee) PrintInfo() {
	// 调用组合的 Person 结构体的 PrintInfo() 方法输出员工的信息
	fmt.Printf("Employee Name :%s, Employee Age: %d, Employee ID: %d\n", e.Person.Name, e.Person.Age, e.EmployeeID)
}

func main() {
	person := Person{Name: "John", Age: 30}
	employee := Employee{EmployeeID: 1001, Person: person}
	employee.PrintInfo()

}


