package main

import "fmt"
/* 合并区间：
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，
然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
如果没有重叠，则将当前区间添加到切片中。 */

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	// 按照区间的起始位置进行排序
	sort(intervals)
	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1] = []int{merged[len(merged)-1][0], max(intervals[i][1], merged[len(merged)-1][1])}
		} else {
			merged = append(merged, intervals[i])
		}
	}
	fmt.Println(merged)
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sort(intervals [][]int) {
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(intervals)
	intervals = [][]int{{1, 4}, {4, 5}}
	merge(intervals)
	intervals = [][]int{{4, 7}, {1, 4}}
	merge(intervals)
}

