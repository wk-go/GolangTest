package main

import "fmt"

type Student struct {
	Age int
}

func main() {
	// for range 中使用指针会怎样？
	a := []Student{{Age: 3}, {Age: 2}, {Age: 1}}
	b := make([]*Student, 3)
	for k, v := range a {
		fmt.Println(k, v, &v)
		b[k] = &v //这样做b切片里只会存储最后一个值
	}
	fmt.Println(b[0], b[1], b[2])
}
