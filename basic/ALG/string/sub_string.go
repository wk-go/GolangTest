package main

//给定一个字符串，请你找出其中不含有重复字符的最长子串的长度
import "fmt"

func main() {
	strSlice := []string{
		"abcbacdacef",
		"ababab",
		"aaa",
		"abcdefg",
	}
	for i := range strSlice {
		length := subString(strSlice[i])
		fmt.Printf("\"%s\" sub string length:%d\n", strSlice[i], length)
	}
}

func subString(s string) int {
	tmp := map[byte]int{}
	length := len(s)
	right, maxLen := 0, 0
	for left := 0; left < length; left++ {
		if left > 0 {
			delete(tmp, s[left-1])
		}

		for right < length && tmp[s[right]] == 0 {
			tmp[s[right]]++
			right++
		}
		if t := right - left; t > maxLen {
			maxLen = t
		}
	}
	return maxLen
}
