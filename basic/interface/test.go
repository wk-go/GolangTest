package main

import "fmt"

// 接口与普通类型比较

func main() {
	var param1 interface{} = 1
	var param2 int = 1
	if param1 == param2 {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not Equal!!")
	}
}
