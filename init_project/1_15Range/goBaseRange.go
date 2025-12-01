package main

import (
	"fmt"
	"reflect"
	"time"
)
func main() {
	// -----------------Range 迭代字符串-------------
	var str1 string = "abc123efg"
	for index , value := range str1 {
		fmt.Printf("str1[%d] = %c\n", index, str1[index])
		fmt.Printf("str1[%d] = %c\n", index, value)
	}

	var str2 string = "Hello, 世界"
	for index, value := range str2 {
		fmt.Printf("含中文字符串 str2[%d] = %c\n", index, str2[index])
		fmt.Printf("含中文字符串 str2[%d] = %c\n", index, value)
	}
	str3 := "a1中文"
    for index, value := range str3 {
		// 这里的value是rune类型，表示一个Unicode码点，而非byte类型,输出的d%
        fmt.Printf("str3 -- index:%d, index value:%d\n", index, str3[index])
        fmt.Printf("str3 -- index:%d, range value:%d\n", index, value)
    }

	runesFromStr2 := []rune(str2)
    bytesFromStr2 := []byte(str2)
    fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))
    fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2))

	// -----------------Range 迭代数组和切片--------------------------
	var array [3]int = [3]int{1,2,3}
	var slice []int = []int{4,5,6}
	// 方法1：只拿到数组的下标索引
	for index := range array {
		fmt.Printf("array -- index=%d value=%d \n", index, array[index])
	}
	for index := range slice {
        fmt.Printf("slice -- index=%d value=%d \n", index, slice[index])
    }
	fmt.Println("-")
	// 方法2：拿到数组和切片的下标索引和值
	for index, value := range array {
		fmt.Printf("array -- index=%d index value=%d \n", index, array[index])
        fmt.Printf("array -- index=%d range value=%d \n", index, value)
	}
	for index, value := range slice {
		fmt.Printf("slice -- index=%d index value=%d \n", index, slice[index])
        fmt.Printf("slice -- index=%d range value=%d \n", index, value)
	}
	// -----------------Range 遍历多维数组------------------
	array2 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
    slice2 := [][]int{{1, 2}, {3}}
    // 只拿到行的索引
    for index := range array2 {
        // array2[index]类型是一维数组
        fmt.Println(reflect.TypeOf(array2[index]))
        fmt.Printf("array2 -- index=%d, value=%v\n", index, array2[index])
    }

    for index := range slice2 {
        // slice[index]类型是一维数组
        fmt.Println(reflect.TypeOf(slice2[index]))
        fmt.Printf("slice -- index=%d, value=%v\n", index, slice2[index])
    }

    // 拿到行索引和该行的数据
    fmt.Println("print array2 element")
    for row_index, row_value := range array2 {
        fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
    }

    fmt.Println("print array2 slice2")
    for row_index, row_value := range slice2 {
        fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
    }

    // 双重遍历，拿到每个元素的值
    for row_index, row_value := range array2 {
        for col_index, col_value := range row_value {
            fmt.Printf("array[%d][%d]=%d ", row_index, col_index, col_value)
        }
        fmt.Println()
    }
    for row_index, row_value := range slice2 {
        for col_index, col_value := range row_value {
            fmt.Printf("slice[%d][%d]=%d ", row_index, col_index, col_value)
        }
        fmt.Println()
    }
	// -----------------Range 迭代通道channel----------------
	ch := make(chan int, 10)
    go addData(ch)
    for i := range ch {
        fmt.Println("通道数据：", i) // 1秒钟打印一次数据
    }
	// -----------------Range 迭代Map--------------------
	hash := map[string]int{
        "a": 1,
        "f": 2,
        "z": 3,
        "c": 4,
    }

    for key := range hash {
        fmt.Printf("key=%s, value=%d\n", key, hash[key])
    }

    for key, value := range hash {
        fmt.Printf("key=%s, value=%d\n", key, value)
    }
}

func addData(ch chan int) {
    size := cap(ch)
    for i := 0; i < size; i++ {
        ch <- i
        time.Sleep(1 * time.Second) // 1秒钟发送一次数据
    }
    close(ch)
}