package main

import (
	"fmt"
)
/* 题目：查找字符串数组中的最长公共前缀  */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs)) // Output: fl
	strs = []string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(strs)) // Output: ""
}