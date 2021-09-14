package main

import "fmt"

// 浮点数的陷阱
func main() {
	x := 0.7 + 0.2
	fmt.Printf("%.40f\n", x)

	x = 0.1 + 0.2
	fmt.Printf("%.40f\n", x)
}
