package main

import "fmt"

func main() {
	var sum float64
	for i := 2; i < 22; i++ {
		fmt.Printf("f(%d)/f(%d)=%d/%d,,\n", i+1, i, fibonacci(i+1), fibonacci(i))
		sum += float64(fibonacci(i+1)) / float64(fibonacci(i))
	}
	fmt.Println("result:", sum)
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	n1, n2 := 0, 1
	tmp := 0
	for i := 2; i <= n; i++ {
		tmp = n2
		n2 = n1 + n2
		n1 = tmp
	}
	return n2
}
