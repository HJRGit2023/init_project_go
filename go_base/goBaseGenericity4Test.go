package main

import "fmt" 

func Filter[T any](slice []T,predicate func(T)bool) []T {
	// 实现过滤逻辑
	result := make([]T, 0)
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

// 你的实现代码在这里
func Map[T any, U any](slice []T, mapper func(T) U)[]U {
   // 实现映射逻辑
   result := make([]U, 0)
   for _, v := range slice {
       result = append(result, mapper(v))
   }
   return result
}

func main() {
	// 测试用例
	numbers :=[]int{1,2,3,4,5,6}
	even := Filter(numbers,func(n int)bool{
	return n%2==0
	})
	fmt.Println(even)  // 应该输出: [2 4 6]

	numbers2 :=[]int{1,2,3,4,5}
	strings := Map(numbers2,func(n int)string{
	return fmt.Sprintf("Number: %d", n)
	})
	fmt.Println(strings)
}