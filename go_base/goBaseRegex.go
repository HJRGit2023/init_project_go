package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 检查字符串是否匹配正则表达式
	pattern := `^[a-zA-Z0-9]+$`
	regex := regexp.MustCompile(pattern)
	str := "Hello 123"
	if regex.MatchString(str) {
		fmt.Println(str,"字符串匹配正则表达式")
	} else {
		fmt.Println(str,"字符串不匹配正则表达式")
	}
	str1 := "Hello123"
	if regex.MatchString(str1) {
		fmt.Println(str1,"字符串匹配正则表达式")
	} else {
		fmt.Println(str1,"字符串不匹配正则表达式")
	}

	// 查找匹配的字符串
	pattern1 := `\d+`
	regex1 := regexp.MustCompile(pattern1)
	str2 := "我有 3 个苹果和 5 个香蕉"
	matches := regex1.FindAllString(str2, -1)
	fmt.Println("找到的数字：", matches)
	// 替换匹配后的字符串
	pattern2 := `\s+` // 匹配空白字符（包括空格、制表符、换行符等）。
	regex2 := regexp.MustCompile(pattern2)
	str3 := "Hello, world!"
	results := regex2.ReplaceAllString(str3, "*")
	fmt.Println("替换后的字符串：", results)
	// 分割字符串
	pattern3 := `,` // 匹配逗号。
	regex3 := regexp.MustCompile(pattern3)
	str4 := "apple,banana,orange"
	results1 := regex3.Split(str4, -1)
	fmt.Println("分割后的字符串：", results1)

}