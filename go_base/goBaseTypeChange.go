package main

import "fmt"
import "strconv"

// 定义一个接口 Writer
	type Writer interface {
		Write([]byte)(int, error)
	}
	// 实现 Writer 接口的结构体 StringWriter
	type StringWriter struct {
		str string
	}
	// func struct(结构体) interface(接口) 返回类型int和err信息
	func (sw *StringWriter) Write(p []byte) (int,error) {
		sw.str += string(p)
		return len(p), nil
	}
	// 空接口 interface{} 可以持有任何类型的值。在实际应用中，空接口经常被用来处理多种类型的值。
	func printValue(v interface{}) {
		switch v := v.(type) { // v.(type) 语法用于类型断言，判断 v 的实际类型，v接口类型转为指定类型type
		case int:
			fmt.Println("Integer:", v)
		case string:
			fmt.Println("String:", v)
		default:
			fmt.Println("Unknown type")
		}
	}

func main() {
	var sum int = 17
	var count int = 5
	var mean float32
	// sum值int类型转为 float32类型：float32(sum)
	mean = float32(sum)/float32(count)
	fmt.Printf("mean's value is %f\n", mean)

	var str string = "123"
	var num,num1 int
	// str值string类型转为 int类型：num, err := strconv.Atoi(str)
	num, _ = strconv.Atoi(str) // _忽略错误处理
	fmt.Printf("num's value is %d\n", num)
	num1, err := strconv.Atoi(str)
	if err !=nil {
		fmt.Println("转换错误:", err)
	}else{
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num1)
	}

	str1 := strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num, str1)

	// 字符串转为浮点数：num2, err := strconv.ParseFloat(str, 64)
	str2 := "3.14"
	num2, err := strconv.ParseFloat(str2, 64)
	if err !=nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为浮点数为：%f\n", str2, num2)
	}
	// 浮点数转为字符串：
    str3 := strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转换为字符串为：'%s'\n", num2, str3)

	// 接口类型变量
	var i interface{} = "Hello, world!"
	str4 , ok := i.(string)
	if ok {
		fmt.Printf("接口类型变量 i 的值是字符串 '%s'\n", str4)
	} else {
		fmt.Println("接口类型变量 i 的值不是字符串")
	}
	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = &StringWriter{}
	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(*StringWriter)
	// 修改 StringWriter 的字段
	sw.str = "Hello, world!"
	// 打印 StringWriter 的字段值
	fmt.Printf("StringWriter 的字段值是：'%s'\n", sw.str)

	printValue(42)
    printValue("hello")
    printValue(3.14)
}