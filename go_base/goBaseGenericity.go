package main

import "fmt"

// 数字类型约束
type Number interface{
	int|int8|int16|int32|int64| 
	uint|uint8|uint16|uint32|uint64| 
	float32|float64
}
// 自定义约束
type Stringer interface {
	String() string
}

func printString[T Stringer](value T) {
	fmt.Println(value.String())
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
	// 泛型
	// 使用示例
	PrintAny(42)        // Value: 42, Type: int
	PrintAny("hello")   // Value: hello, Type: string
	PrintAny(3.14)      // Value: 3.14, Type: float64

	// 使用示例
	numbers :=[]int{1,2,3,4,5}
	fmt.Println("3的索引位置是：",FindIndex(numbers,3))  // 输出: 2

	names :=[]string{"Alice","Bob","Charlie"}
	fmt.Println("Bob的索引位置是：",FindIndex(names,"Bob"))  // 输出: 1

	// 使用示例
	fmt.Println("两int数之和为：", Add(10,20))        // 输出: 30
	fmt.Println("两float数之和为：", Add(3.14,2.71))    // 输出: 5.85

	// 自定义约束 使用示例
	person := Person{Name:"Alice", Age:25}
	printString(person)
}

func Add[T Number](a, b T) T {
	return a + b
}

// 函数名[T 约束](参数 T) 返回值类型
// func Max[T comparable](a, b T) T {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// any 约束 : any 是空接口 interface{} 的别名，表示任何类型都可以。
func PrintAny[T any](value T){
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}
// comparable约束 :comparable 表示类型支持 == 和 != 操作符。
func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}




