package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"io/ioutil"
)

func main() {
	file, err := os.Create("Write1.txt")
	if err!= nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// 方式1：写入字符串
	file.WriteString("直接写入字符串\n")
	// 方式2：写入字节切片
	file.Write([]byte("写入字节数组\n"))
	data := []byte("写入字节切片\n")
	file.Write(data)
	// 方式3：使用fmt.Fprintf格式化写入
	fmt.Fprintf(file,"格式化写入%d\n", 123)
	// 方式4：逐行写入
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, "逐行写入", i)
	}
	writer.Flush() // 调用 writer.Flush() 确保所有数据都被写入文件。
	log.Println("写入完成")
	// 一次性写入
	content :=[]byte("一次性写入，Hello, World!")
	// 0644表示文件权限 表示文件所有者可以读写，其他用户只能读取。
	err = ioutil.WriteFile("Write1.txt", content,0644) 
	if err !=nil{
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("一次性写入完成")
	// 追加写入
	file, err = os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err!= nil {
		log.Fatalln(err)
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	if _, err := file.WriteString("追加写入\n"); err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("追加写入完成")
	// 手动关闭文件
	if err = file.Close(); err != nil {
		fmt.Println("Error closing file:", err)
		return
	}
	// 删除文件
	err = os.Remove("output.txt")
	if err != nil {
		fmt.Println("Error removing file:", err)
		return
	}
	fmt.Println("文件删除完成")
	// 获取文件信息
	fileInfo, err := os.Stat("Write1.txt")
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fmt.Println("文件信息：", fileInfo)
	fmt.Println("文件名称：", fileInfo.Name())
	fmt.Println("文件大小：", fileInfo.Size())
	fmt.Println("文件权限：", fileInfo.Mode())
	fmt.Println("文件创建时间：", fileInfo.ModTime())
	fmt.Println("文件修改时间：", fileInfo.ModTime())
	fmt.Println("文件是否是目录：", fileInfo.IsDir())
	fmt.Println("文件是否可读：", fileInfo.Mode().Perm()&os.ModePerm)
	fmt.Println("文件是否可写：", fileInfo.Mode().Perm()&os.ModePerm)
	fmt.Println("文件是否可执行：", fileInfo.Mode().Perm()&os.ModePerm)
	fmt.Println("文件是否可链接：", fileInfo.Mode().Perm()&os.ModePerm)
	fmt.Println("文件是否可删除：", fileInfo.Mode().Perm()&os.ModePerm)
	fmt.Println("文件是否可重命名：", fileInfo.Mode().Perm()&os.ModePerm)

	// 判断文件是否存在
	if _, err := os.Stat("Write1.txt"); os.IsNotExist(err) {
		fmt.Println("Write1.txt文件不存在")
	} else {
		fmt.Println("Write1.txt文件存在")
	}
	// 重命名文件，先手动关闭文件，再重命名文件。
	file2, err := os.OpenFile("test.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// 进行一些文件操作...
	// 关闭文件
	file2.Close()
	err = os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}
	fmt.Println("文件重命名完成")
}