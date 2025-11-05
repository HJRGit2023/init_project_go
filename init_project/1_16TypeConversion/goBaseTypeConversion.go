package main

import "fmt"
import "strconv"
// 接口类型 Supplier 转换为 DigitSupplier结构体类型
type Supplier interface {
    Get() string
}

type DigitSupplier struct {
    value int
}

func (i *DigitSupplier) Get() string {
    return fmt.Sprintf("%d", i.value)
}
// 两个结构体(名称不一致) 字段名和类型相同，可以互相转换
type SameFieldA struct {
    name  string
    value int
}

type SameFieldB struct {
    name  string
    value int
}

func (s *SameFieldB) getValue() int {
    return s.value
}

func main() {
	// ----------------------- 数字类型转换 ----------
	var i int32 = 17
    var b byte = 5
    var f float32
    
    // 数字类型可以直接强转
    f = float32(i) / float32(b)
    fmt.Printf("f 的值为: %f\n", f)
    
    // 当int32类型强转成byte时，高位被直接舍弃
    var i2 int32 = 256
    var b2 byte = byte(i2)
    fmt.Printf("b2 的值为: %d\n", b2) // 输出 0, 因为 256 超过了 byte 类型最大值 255
   // ---------- 字符串类型转换 ----------
   str := "hello, 123, 你好"
    var bytes []byte = []byte(str) // byte 8位
    var runes []rune = []rune(str) // rune 32位 等价于 int32
    fmt.Printf("bytes 的值为: %v \n", bytes)
    fmt.Printf("runes 的值为: %v \n", runes)

    str2 := string(bytes) // byte 转字符串 不会丢失信息
    str3 := string(runes) // rune 转字符串 不会丢失信息
    fmt.Printf("str2 的值为: %v \n", str2)
    fmt.Printf("str3 的值为: %v \n", str3)
	// -----------数字和字符串的转换------------------
	str4 := "123"
    num, err := strconv.Atoi(str4)
    if err != nil {
        panic(err)
    }
    fmt.Printf("字符串转换为int: %d \n", num)
    str5 := strconv.Itoa(num)
    fmt.Printf("int转换为字符串: %s \n", str5)

    ui64, err := strconv.ParseUint(str4, 10, 32)
    fmt.Printf("字符串转换为uint64: %d \n", num)

    str6 := strconv.FormatUint(ui64, 2)
    fmt.Printf("uint64转换为字符串: %s \n", str6)
	//------------------接口类型转换----------------
	var ii interface{} = 3
    aa, ok := ii.(int)
    if ok {
        fmt.Printf("'%d' is a int \n", aa)
    } else {
        fmt.Println("conversion failed")
    }
	// 上面的类型转换 Switch 语句也可以实现
	switch vv := ii.(type) {
    case int:
        fmt.Println("ii is a int", vv)
    case string:
        fmt.Println("ii is a string", vv)
    default:
        fmt.Println("ii is unknown type", vv)
    }
	// -------------------结构体类型转换---------------------
	var aaa Supplier = &DigitSupplier{234}
	fmt.Println(aaa.Get()) // 输出 234

	bbb, ok := aaa.(*DigitSupplier)
	fmt.Println(bbb, ok) // 输出 &{234} true
	if ok {
		fmt.Println(bbb.Get()) // 输出 234
	} else {
		fmt.Println("conversion failed")
	}
	// 两个结构体(名称不一致) 字段名和类型相同，可以互相转换
	aaaa := SameFieldA{
        name:  "aaaa",
        value: 1,
    }

    bbbb := SameFieldB(aaaa)
    fmt.Printf("conver SameFieldA to SameFieldB, value is : %d \n", bbbb.getValue())
    
    // 只能结构体类型实例之间相互转换，指针  和 结构体 不可以相互转换
    // var c interface{} = &a
    // _, ok := c.(*SameFieldB)
    // fmt.Printf("c is *SameFieldB: %v \n", ok)
}