package main

import (
	"fmt"
	"os"
	"log"
	"io"
)

func main() {
	// 创建单个目录
	err :=os.Mkdir("d:\\web3\\Go_Workspace\\testDir", 0755)
	if err!= nil {
		fmt.Println(err)
	} else {
		fmt.Println("创建单个目录成功")
	}
	// 创建多个目录
	err = os.MkdirAll("d:\\web3\\Go_Workspace\\testDir\\testDir2\\testDir3", 0755)
	if err!= nil {
		fmt.Println(err)
	} else {
		fmt.Println("创建多个目录成功")
	}
	// 读取目录内容
	entries, err := os.ReadDir("d:\\web3\\Go_Workspace\\testDir")
	if err != nil {
		fmt.Println(err)
	}
	//  “左对齐、宽度 20 的文件名 + 右对齐、宽度 8 的文件大小 + 文件修改时间” 的格式
	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Printf("%-20s %8d %v\n", 
			entry.Name(), 
			info.Size(), 
			info.ModTime().Format("2006-01-02 15:04:05"))
	}

	 // 删除空目录
	err = os.Remove("d:\\web3\\Go_Workspace\\testDir\\testDir2\\testDir3")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("删除空目录成功")
	}
	
	// 递归删除目录及其内容
	err = os.RemoveAll("d:\\web3\\Go_Workspace\\testDir")
	if err != nil {
			log.Fatal(err)
	} else {
		fmt.Println("递归删除目录及其内容成功")
	}

	// 文件复制：打开文件source.txt
	srcFile, err := os.Open("source.txt")
	if err !=nil{
		log.Fatal(err)
	}

	// 文件复制：创建文件 destination.txt
	dstFile, err := os.Create("destination.txt")
	if err !=nil{
		log.Fatal(err)
	}
	defer dstFile.Close()
	// 文件复制：复制文件内容
	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err !=nil{
		log.Fatal(err)
	}
	log.Printf("复制完成，共复制 %d 字节", bytesCopied)

	// 文件追加：打开文件 destination.txt
	dstFile, err = os.OpenFile("destination.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err !=nil{
		log.Fatal(err)
	}
	// 文件追加：写入内容
	_, err = dstFile.WriteString("追加内容")
	if err !=nil{
		log.Fatal(err)
	}
	log.Println("追加完成")

	// 创建临时文件
	tempFile, err := os.CreateTemp("d:\\web3\\Go_Workspace\\", "tempFile")
	if err !=nil{
		log.Fatal(err)
	}
	defer func() {
		tempFile.Close() // 关闭临时文件
		os.Remove(tempFile.Name()) // 删除临时文件
	}()
	fmt.Println("临时文件名：", tempFile.Name())
	// 写入内容
	_, err = tempFile.WriteString("临时文件内容")
	if err !=nil{
		log.Fatal(err)
	}
	log.Println("创建临时文件成功")
	// 创建临时目录
	tempDir, err := os.MkdirTemp("d:\\web3\\Go_Workspace\\", "tempDir-*")
	if err !=nil{
		log.Fatal(err)
	}
	log.Println("创建临时目录成功")
	defer os.RemoveAll(tempDir) // 删除临时目录

}