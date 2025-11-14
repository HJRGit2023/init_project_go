// 冒泡排序
package main

import "fmt"	

func main() {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	fmt.Println("Before sorting:", arr)
	bubbleSort(arr)
	fmt.Println("After sorting:", arr)
}

bubbleSort := func(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("After sorting:", arr)
}