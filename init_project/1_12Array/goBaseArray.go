package main

import "fmt"

type Custom struct {
    i int
}

var carr [5]*Custom = [5]*Custom{
    {6},
    {7},
    {8},
    {9},
    {10},
}

func main() {
	// --------------------------数组声明5个方式----------
	// 方式一：仅声明数组，未初始化
	var arr [5]int
	fmt.Println(arr)
	var marr [2]map[string]string
	fmt.Println(marr)
	// map[string]string{} 等价于 map[string]string(nil)
	// map的零值是 nil，虽然打印出来是 非空值，但实际上是 nil

	// 方式二：声明以及初始化数组
	var b [5]int = [5]int{1,2,3,4,5}
	fmt.Println("b = ", b)

	// 方式三：类型推导
	var c = [5]string{"c1", "c2", "c3", "c4", "c5"}
	fmt.Println("c = ", c)

	d := [3]int{1, 2, 3}
	fmt.Println("d = ", d)
	// 方式四：数组长度自动推导
	autoLen := [...]int{1,2,3,4}
	fmt.Println("autoLen = ", autoLen)
    // 方式五：使用数组下标赋值
	positionInit := [5]string{1: "positionInit1", 3: "positionInit3"}
	fmt.Println("positionInit = ", positionInit)
	// 初始化时，元素个数不能超过数组声明的长度
    //overLen := [2]int{1, 2, 3}

	// -----------------------数组访问-------------------------
	// 方式一：通过下标访问
	fmt.Println("c[0] = ", c[0])
	// 方式二：通过for循环遍历数组
	for i := 0; i < len(c); i++ {
		fmt.Println("for循环，使用len(c) 下标 = ", i, " 值 =", c[i])
	}
	// 方式三：通过range遍历数组
	for i, v := range c {
		fmt.Println("index = ", i, "value = ", v)
	}
	// 方式四：通过指针访问数组元素	
	var e [5]int
	p := &e
	p[0] = 10
	fmt.Println("e = ", e)
	// -----------------------多维数组-----------------
	// 二维数组
	mula := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("mula = ", mula)
    // 三维数组
	mula3 := [3][2][2]int{
        {{0, 1}, {2, 3}},
        {{4, 5}, {6, 7}},
        {{8, 9}, {10, 11}},
    }
	fmt.Println("mula3 = ", mula3)
	// 也可以省略各个位置的初始化,在后续代码中赋值
    mulc := [3][3][3]int{}
    mulc[2][2][1] = 5
    mulc[1][2][1] = 4
    fmt.Println("mulc = ", mulc)
	// -----------------------多维数组访问------------------
	layer1 := mula3[0]
    layer2 := mula3[0][1]
    element := mula3[0][1][1]
    fmt.Println("layer1 = ", layer1)
    fmt.Println("layer2 = ", layer2)
    fmt.Println("layer2[1] = ", element)

    // 多维数组遍历时，需要使用嵌套for循环遍历
    for i, v := range mula3 {
        fmt.Println("index = ", i, "value = ", v)
        for j, inner := range v {
            fmt.Println("inner, index = ", j, "value = ", inner)
        }
    }

	// 数组作为函数参数
	printFuncParamPointer(carr)
    arr1 := [5]int{1, 2, 3, 4, 5}
	receiveArray(arr1)
	fmt.Println("unchanged ,main func, after modify, arr1 = ", arr1)
	receiveArrayPointer(&arr1)
	fmt.Println("use pointer changed ,main func, after modify, arr1 = ", arr1)
}

func receiveArray(param [5]int) {
    fmt.Println("in receiveArray func, before modify, param = ", param)
    param[1] = -5
    fmt.Println("in receiveArray func, after modify, param = ", param)
}

func receiveArrayPointer(param *[5]int) {
    fmt.Println("in receiveArrayPointer func, before modify, param = ", param)
    param[1] = -5
    fmt.Println("in receiveArrayPointer func, after modify, param = ", param)
}

func printFuncParamPointer(param [5]*Custom) {
    for i := range param {
        fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
    }
}
