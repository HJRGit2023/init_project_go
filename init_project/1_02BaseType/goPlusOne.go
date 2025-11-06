package main

import "fmt"
/* 加一  ：
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。*/
func plusOne(nums []int) {
	if len(nums) == 0 {
		return
	}
	nums[len(nums)-1] += 1 // 结尾加1
	for i := len(nums) - 1; i >= 0; i-- { // 从结尾开始向前遍历
		if nums[i] == 10 {
			nums[i] = 0 // 进位
			if i == 0 { // 若是最高位，则开头加1
				nums = append([]int{1}, nums...) // 开头加1
			} else {
				nums[i-1] += 1 // 进位
			}
		} else {
			break
		}
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{9, 9, 9}
	plusOne(nums)
	nums = []int{1,2,3}
	plusOne(nums)
	nums = []int{4,3,2,1}
	plusOne(nums)
	nums = []int{0}
	plusOne(nums)
	nums = []int{9}
	plusOne(nums)


}