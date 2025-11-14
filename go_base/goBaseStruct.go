package main

import "fmt"

// Book struct 结构体定义需要使用 type 和 struct 语句。
// struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。
type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	// struct 可以用
	// 创建一个新的结构体
    fmt.Println(Books{"Go 语言","www.runoob.com","Go 语言教程",6495407})
	// 也可以使用 key => value 格式
    fmt.Println(Books{title:"Go 语言", author:"www.runoob.com", subject:"Go 语言教程", book_id:6495407})
	// 忽略的字段为 0 或 空
   fmt.Println(Books{title:"Go 语言", author:"www.runoob.com"})

   var book1 Books /* 声明 book1 为 Books 类型 */
   var book2 Books /* 声明 book2 为 Books 类型 */
   book1.title = "Go 语言" /*要访问结构体成员，需要使用点号 . 操作符，格式为： 结构体.成员名"*/
   book1.author = "www.runoob.com"
   book1.subject = "Go 语言教程"
   book1.book_id = 6495407
   fmt.Println(book1)
   /* 打印 Book1 信息 */
fmt.Printf("Book 1 title : %s\n", book1.title)
fmt.Printf("Book 1 author : %s\n", book1.author)
fmt.Printf("Book 1 subject : %s\n", book1.subject)
fmt.Printf("Book 1 book_id : %d\n", book1.book_id)

   book2.title = "Python 教程"
   book2.author = "www.runoob.com"
   book2.subject = "Python 语言教程"
   book2.book_id = 6495408
   fmt.Println(book2)
   /* 打印 Book2 信息 */
fmt.Printf("Book 2 title : %s\n", book2.title)
fmt.Printf("Book 2 author : %s\n", book2.author)
fmt.Printf("Book 2 subject : %s\n", book2.subject)
fmt.Printf("Book 2 book_id : %d\n", book2.book_id)
/* 调用 printBook 函数打印 Book1 和 Book2 信息 */
printBook(book1)
printBook(book2)
/* 调用 printBookPointer 函数  结构体指针访问结构体成员  打印 Book1 和 Book2 信息 */
printBookPointer(&book1)
printBookPointer(&book2)
}

func printBook( book Books ) {
   fmt.Printf( "函数Book title : %s\n", book.title)
   fmt.Printf( "函数Book author : %s\n", book.author)
   fmt.Printf( "函数Book subject : %s\n", book.subject)
   fmt.Printf( "函数Book book_id : %d\n", book.book_id)
}

func printBookPointer(book *Books) {
	fmt.Printf( "函数 结构体指针访问结构体成员 Book title : %s\n", book.title)
   fmt.Printf( "函数 结构体指针访问结构体成员 Book author : %s\n", book.author)
   fmt.Printf( "函数 结构体指针访问结构体成员 Book subject : %s\n", book.subject)
   fmt.Printf( "函数 结构体指针访问结构体成员 Book book_id : %d\n", book.book_id)
}

