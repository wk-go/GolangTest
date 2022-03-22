package main

import "fmt"

type Student struct {
	Age int
}

func main() {
	// for range 中使用指针会怎样？
	a := map[string]Student{"Sam": {Age: 3}, "Peter": {Age: 2}, "Alice": {Age: 1}}
	b := make([]*Student, 0, 3)
	for k, v := range a {
		fmt.Println(k, v, &v)
		b = append(b, &v) //这样做b切片里只会存储最后一个值
	}
	fmt.Printf("%#v\n", b)
	fmt.Println(b[0], b[1], b[2])

	a1 := map[string]*Student{"Sam": {Age: 3}, "Peter": {Age: 2}, "Alice": {Age: 1}}
	b1 := make([]*Student, 0, 3)
	for k, v := range a1 {
		fmt.Println(k, v, &v)
		b1 = append(b1, v) //这样做b切片里只会存储最后一个值
	}
	fmt.Printf("%#v\n", b)
	fmt.Println(b1[0], b1[1], b1[2])
}
