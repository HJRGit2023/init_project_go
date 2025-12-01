package main

import (
	// "fmt" // 引入了需要使用 "fmt" imported and not used
	"fmt"
	"bufio"
	"os"
	"log"
	"io/ioutil"
)
/* 1 创建文件
2 打开文件
3 写入文件
4 读取文件
*/
func main() {
	// 1.1创建文件，如果文件已存在会被截断（清空）
	file, err := os.Create("test.txt")
	if err!=nil {
		log.Fatalln("文件创建失败：", err) // 打印错误信息并退出程序
	}
	defer file.Close() // 确保文件关闭
	log.Println("文件创建成功")
	// 2 打开文件
	file, err = os.OpenFile("test.txt", os.O_RDWR, 0666)
	if err!=nil {
		log.Fatalln("文件打开失败：", err) // 打印错误信息并退出程序
	}
	defer file.Close() // 确保文件关闭
	log.Println("文件打开成功")
	// 3.写入文件
	_, err = file.WriteString("Hello, world!\n")
	if err!=nil {
		log.Fatalln("文件写入失败：", err) // 打印错误信息并退出程序
	}
	log.Println("文件写入成功")
	// 4.读取文件
	// content, err := file.ReadString('\n')
	// 写入文件后，重置文件指针到开头
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalln("Seek file error:", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err !=nil{
		fmt.Println("Error reading file:", err)
	}
	log.Println("文件读取成功")
	// 5.一次性读取整个文件
	content, err := ioutil.ReadFile("test.txt")
	if err !=nil{
		fmt.Println("Error reading file:", err)
	return
	}

	fmt.Println("一次性读取整个文件：",string(content))

}