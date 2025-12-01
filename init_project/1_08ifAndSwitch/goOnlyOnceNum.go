package main

import "fmt"

/* 只出现一次的数字：
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。 */
func main() {
	// 给定一个非空整数数组 除了某个元素只出现一次以外，其余每个元素均出现两次
	var arrnum [9]int = [9]int{1, 2, 3, 1, 2, 3, 4, 4, 5}
	var mapnum = make(map[int]int) // 定义一个map用来统计每个元素出现的次数
	for i :=0; i<len(arrnum); i++ {
		if _, ok := mapnum[arrnum[i]]; ok { // 如果map中已经有这个元素，则次数加1
			mapnum[arrnum[i]]++
		} else {
			mapnum[arrnum[i]] = 1 // 如果map中没有这个元素，则次数设为1
		}
	}

	for key, value := range mapnum {
		if value == 1 {
			fmt.Println("只出现一次的元素为：", key) // 输出只出现一次的元素
		} else {
			fmt.Printf("元素 %d 出现了 %d 次\n", key, value)
		}
	}
	// 输出结果：
	/* 元素 1 出现了 2 次
		元素 2 出现了 2 次
		元素 3 出现了 2 次
		元素 4 出现了 2 次
		只出现一次的元素为： 5 */
}