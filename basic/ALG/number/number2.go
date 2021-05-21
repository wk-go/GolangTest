package main

import "fmt"

//

func main() {
	nums := []int{23, 2, 3, 5, 33, 66, 22, 99, 20, 30, 56, 87}
	lower, upper := 19, 30
	count := getSectionCount(nums, lower, upper)
	fmt.Printf("count:%d", count)
}

func getSectionCount(nums []int, lower, upper int) int {
	sum := 0
	length := len(nums)
	count := 0
	for left := 0; left < length; left++ {
		sum = nums[left]
		if sum >= lower && sum <= upper {
			count++
		}
		for right := left + 1; right < length; right++ {
			sum += nums[right]
			if sum >= lower && sum <= upper {
				count++
			}
		}
	}
	return count
}
