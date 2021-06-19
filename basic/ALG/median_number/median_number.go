package main

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数
import "fmt"

func main() {
	arr1 := []int{1, 5, 8, 9, 25, 60}
	arr2 := []int{2, 6, 20, 30, 40, 51, 63, 72}
	num := findMedianNumber(arr1, arr2)
	fmt.Println("result: ", num)
}

//最易于理解的解法
func findMedianNumber(arr1, arr2 []int) float64 {
	var (
		len1   = len(arr1)
		len2   = len(arr2)
		len    = len1 + len2
		value1 int
		value2 int
		i1     int
		i2     int
		i      int
	)

	for i <= len/2 {
		if i1 < len1 && arr1[i1] < arr2[i2] {
			if i == (len/2 - 1) {
				value1 = arr1[i1]
			}
			if i == len/2 {
				value2 = arr1[i1]
			}
			i1++
		} else if i2 < len2 {
			if i == (len/2 - 1) {
				value1 = arr2[i2]
			}
			if i == len/2 {
				value2 = arr2[i2]
			}
			i2++
		}
		i++
	}
	if len%2 == 0 {
		return (float64(value1) + float64(value2)) / 2
	}
	return float64(value2)
}
