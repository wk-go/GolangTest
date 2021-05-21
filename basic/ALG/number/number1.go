package main

import (
	"fmt"
	"sort"
)

// 给定一个未排序的整数数组nums，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
func main() {
	nums := []int{10, 23, 4, 5, 6, 35, 33, 66, 9, 7, 11, 15, 13, 12, 22, 24, 25, 14, 8}
	//sort.Ints(nums)
	//fmt.Printf("sorted:%v\n", nums)
	maxLen := scanSequence(nums)
	fmt.Printf("max sequence length:%d\n", maxLen)
}

func scanSequence(nums []int) int {
	sort.Ints(nums)
	maxLen, l := 0, 1
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i] == (nums[i-1] + 1) {
			l++
		} else {
			l = 1
		}
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}
