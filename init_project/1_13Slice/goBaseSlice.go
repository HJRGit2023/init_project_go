package main

import "fmt"

func main() {
	// ------------切片 声明方式----------------
	// 方式一：声明并初始化一个空的切片
	var s1 []int = []int{}
	fmt.Println("s1 = ", s1) // []

	// 方式二：类型推导，并初始化一个空的切片
	s2 := []int{}
	fmt.Println("s2 = ", s2) // []

	// 方式三：类型推导，与方式二等价
	var s3 = []int{} 
	fmt.Println("s3 = ", s3) // []

	// 方式4，与方式1、2、3 等价，可以在大括号中定义切片初始元素
	s4 := []int{1, 2, 3, 4, 5}
	fmt.Println("s4 = ", s4) // [1 2 3 4 5]

	// 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	s5 := make([]int, 0)
	fmt.Println("s5 = ", s5) // []

	// 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为5，容量为10
	s6 := make([]int, 5, 10)
	fmt.Println("s6 = ", s6) // [0 0 0 0 0]


	// 方式7：声明一个切片，并用字面量语法初始化
	s7 := []int{1: 10, 3: 30, 5: 50}
	fmt.Println("s7 = ", s7) // [0 10 0 30 50]

	// 方式8，引用一个数组，初始化切片
	arr := [5]int{6,5,4,3,2}
	// 从数组下标2开始，直到数组的最后一个元素
	s8 := arr[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s9 := arr[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s10 := arr[:2]
	fmt.Println(s8)
    fmt.Println(s9)
    fmt.Println(s10)
    

	// ------------切片 基本操作----------------------
	// 1.获取切片长度
	fmt.Println("s6 len = ", len(s6)) // 5

	// 2.获取切片容量
	fmt.Println("s6 cap = ", cap(s6)) // 5

	// 3.获取切片元素
	fmt.Println("s6[0] = ", s6[0]) // 1
	fmt.Println("s6[4] = ", s6[4]) // 5

	// 4.修改切片元素
	s6[0] = 100
	fmt.Println("modified s6 =", s6) // [100 2 3 4 5]

	// 当切片是基于  同一个数组指针  创建出来时，修改 数组中的值时，同样 会影响 到这些切片
	arr[0] = 9
    arr[1] = 8
    arr[2] = 7
    fmt.Println(s8)
    fmt.Println(s9)
    fmt.Println(s10)
	// ---------------------访问切片元素---------------------
	// 下标访问切片
	e1 := s4[0]
	e2 := s4[1]
	e3 := s4[2]
	fmt.Println("下标访问切片：", s4)
	fmt.Println("e1 = ", e1, "e2 = ", e2, "e3 = ", e3)

	// 向指定位置赋值
	s4[0] = 10
	s4[1] = 9
	s4[2] = 8
	fmt.Println("向指定位置赋值：", s4)

	// range迭代访问切片
	for i, v := range s4 {
		fmt.Println("after modify, s1[%d] = %d", i, v)
	}
	// ---------------------切片拼接 append()---------------
	s11 := []int{}
	s11 = append(s4) // 错误，append()方法不存在
	s11 = append(s4, 100)
	s11 = append(s4, 101, 102, 103)
	fmt.Println("切片拼接：s11 = ", s11)
	// 向切片指定位置 追加元素
	/* append 方法后的 ... 是变长参数（可变参数）语法，
	用于将切片 “展开” 为多个  独立元素，
	以便 append 可以  逐个接收  这些元素  并添加  到目标切片中 */
	s11 = append(s11[:2], append([]int{3}, s11[2:]...)...)
	fmt.Println("向切片指定位置追加元素：s11 = ", s11)
	// 移除指定位置元素代码
	s12 := []int{1,2,3,5,4}
	s12 = append(s12[:3], s12[4:]...)
	fmt.Println("s12 = ", s12)
	// ---------------------切片截取 copy()，切片复制--------------------------
	/* 内置函数 copy() 把某个切片中的所有元素复制到另一个切片，
	复制的长度是它们中最短的切片长度。 */
	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	// 注意：dst2的容量小于3，会导致复制失败,只能复制3位 ，dst2 = [1 2 3]
	fmt.Println("before copy, dst2 = ", dst2) 

	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("after copy, src1 = ", src1)
	fmt.Println("after copy, dst1 = ", dst1)

	fmt.Println("after copy, src2 = ", src2)
	fmt.Println("after copy, dst2 = ", dst2)

}