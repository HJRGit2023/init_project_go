package main

import "fmt"
import "unsafe"

const (
	
	// 常量还可以用作枚举： 性别 数字 0、1 和 2 分别代表未知性别、女性和男性。
	Unknown = 0
    Female = 1
    Male = 2
	e = "abc"
	f = len(e)
	g = unsafe.Sizeof(e)
)

func main() {
	// 常量声明
   const LENGTH int = 10
   const WIDTH int = 5  
   var area int
   const a, b, c = 1, false, "str" //多重赋值

   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d", area)

   println(a, b, c)  
   println("性别：未知性别%d，女性%d，男性%d", Unknown, Female, Male)
   fmt.Printf("性别2：未知性别%d，女性%d，男性%d\n", Unknown, Female, Male)
	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
   println("const e,f,g = ", e, f, g)
}