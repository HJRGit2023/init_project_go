package main

import "fmt"

func main() {
	// 声明一个包含 2 的幂次方的切片
	var pow = []int{1, 2, 4, 8, 16, 32, 64}
	// 遍历切片并打印 range范围
	for i, v := range pow {
		fmt.Printf("2 的 %d 次方 = %d\n", i, v)
	}
	// Range遍历 string 也可以
	for i, v := range "Hello" {
		fmt.Printf("索引index %d = 字符%c\n", i, v)
	}
	// 创建一个空的 map，key 是 int 类型，value 是 float32 类型
    map1 :=make(map[int]float32)
	// 向 map1 中添加 key-value 对
	map1[1]=1.0
	map1[2]=2.0
	map1[3]=3.0
	map1[4]=4.0
	// 遍历 map 并打印 key-value 对
	for k, v := range map1 {
		fmt.Printf("key=%d, value=%f\n", k, v)
	}
	// 遍历 map1，只读取 key
	for k := range map1 {
		fmt.Printf("key=%d\n", k)
	}

	// 遍历 map1，只读取 value, 注意这里的 _ 是一个占位符，表示忽略这个变量的值
	for _, v := range map1 {
		fmt.Printf("value=%f\n", v)
	}

	ch := make(chan int, 5)
	// 向 ch 中发送 10 个数据
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // 关闭 ch
	// 遍历 ch，打印接收到的数据
	for v := range ch {
		fmt.Printf("接收到的数据：%d\n", v)
	}

	nums :=[]int{2,3,4}

	// 忽略索引
	for _, num :=range nums {
	fmt.Println("忽略索引，打印value:", num)
	}

	// 忽略值
	for i:=range nums {
	fmt.Println("忽略值，打印index:", i)
	}
	//这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
	sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
	//在数组上使用 range 将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("索引为", i, "的元素为", num)
		}
	}
	//range 也可以用在 map 的键值对上。
	kvs := map[string]string{"a" : "apple", "b" : "banana"}
	for k, v := range kvs {
		fmt.Println("Map对应的key:", k, "value:", v)
	}
	//range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go语言" {
		fmt.Println("索引为", i, "的字符为", c)
	}
}