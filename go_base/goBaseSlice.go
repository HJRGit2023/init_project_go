package main

import "fmt"

func main() {
	/*
	var slice1 []type = make([]type, len)
	也可以简写为
	slice1 := make([]type, len)
	*/
	var numbers []int = make([]int, 5)
    printSlice(numbers)
	numbers = make([]int, 3, 5)
	printSlice(numbers)

	var numbers1 []int /* 一个切片在未初始化之前默认为 nil，长度为 0 */
	printSlice(numbers1)
	if numbers1 == nil {
		fmt.Println("numbers1 is nil空的")
	}
	/* 创建切片 */
    numbers2 :=[]int{0,1,2,3,4,5,6,7,8}  
	printSlice(numbers2)
	/* 切片操作  打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers2[1:4] ==", numbers2[1:4])
	/* 切片操作  打印子切片从索引1(包含) 到最后 */
	fmt.Println("numbers2[1:] ==", numbers2[1:])
	/* 默认下限为 0*/
	fmt.Println("numbers2[:3] ==", numbers2[:3])
	/* 默认上限为 len(s)*/
	fmt.Println("numbers2[4:] ==", numbers2[4:])
	numbers3 := make([]int, 0,5)
	fmt.Println("numbers3=",numbers3)
	numbers3 = numbers2[0:2]
	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	fmt.Println("切片从索引  0(包含) 到索引 2(不包含)numbers3 ==", numbers3)
	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	numbers4 := numbers2[2:5]
	fmt.Println("切片从索引 2(包含) 到索引 5(不包含)numbers4=",numbers4)

	var numbers5 []int
    printSlice(numbers5)
    /* 允许追加空切片 */
	numbers5 = append(numbers5, 0)
	printSlice(numbers5)
	/* 向切片添加一个元素 */
	numbers5 = append(numbers5, 1)
	printSlice(numbers5)
	/* 向切片添加多个元素 */
	numbers5 = append(numbers5, 2, 3, 4)
	printSlice(numbers5)
	/* 创建切片 numbers6 是numbers5切片的两倍容量*/
	numbers6 := make([]int, len(numbers5), 2*cap(numbers5))
	printSlice(numbers6)
	/* 拷贝 numbers5 的内容到 numbers6 */
    copy(numbers6,numbers5)
    printSlice(numbers6)   
}

func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}