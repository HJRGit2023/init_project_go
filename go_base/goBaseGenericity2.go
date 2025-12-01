package main

import "fmt"

/* 1.通用工具函数
2. 数字运算函数 */

// 1.1 通用工具函数----交换两个变量
func Swap[T any](a, b T) (T, T) {
	return b, a
}
// 1.2 通用工具函数----判断切片包含元素
func Contains[T comparable](slice []T, target T) bool {
	// 判断切片中是否存在目标值
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// 1.3 通用工具函数----去重函数
func Unique[T comparable] (slice []T) []T {
	seen := make(map[T]bool)
	result := []T {}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// 数字类型约束
type Number interface{
int|int8|int16|int32|int64| 
uint|uint8|uint16|uint32|uint64| 
float32|float64
}

// 2.1 数字运算函数----取最大值函数
func Max[T Number](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	max := slice[0]
	for i := 1; i<len(slice); i++ {
		if (slice[i] > slice[i-1]) {
			max =  slice[i]
		}
	}
	return max
}

func Min[T Number](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	min := slice[0]
	for _, value := range slice[1:] {
		if value < min {
			min = value
		}
	}
	return min
}

// 2.2 数字运算函数----取平均值函数
func Average[T Number](slice []T) float64 {
	if len(slice) == 0 {
		var zero T
		return float64(zero)
	}
	var sum T
	for _, value := range slice {
		sum += value
	}
	return float64(sum) / float64(len(slice))
}

func main() {
	// Swap 示例
	a, b :=10,20
	a, b = Swap(a, b)
	fmt.Printf("a=%d, b=%d\n", a, b)  // 输出: a=20, b=10

	// Contains 示例
	numbers :=[]int{1,2,3,4,5}
	fmt.Println(Contains(numbers,3))  // 输出: true

	// Unique 示例
	duplicates :=[]int{1,2,2,3,4,4,5}
	unique := Unique(duplicates)
	fmt.Println(unique)  // 输出: [1 2 3 4 5]

	ints :=[]int{1,5,3,9,2}
	floats :=[]float64{1.1,5.5,3.3,9.9,2.2}

	fmt.Printf("Max int: %d\n", Max(ints))           // 输出: 9
	fmt.Printf("Min float: %.1f\n", Min(floats))     // 输出: 1.1
	fmt.Printf("Average: %.2f\n", Average(floats))   // 输出: 4.40

}