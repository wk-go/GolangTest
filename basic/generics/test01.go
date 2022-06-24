package main

// 泛型基本用法示例

import "fmt"

func add[T int | float64 | string](a T, b T) T {
	return a + b
}

type addable interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func add2[T addable](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("\nadd:")
	fmt.Println("add(1,2):", add(1, 2))
	fmt.Println("add(1.1,2.2):", add(1.1, 2.2))
	fmt.Println("add(\"Hello \",\"world\"):", add("Hello ", "world"))

	fmt.Println("\nadd2:")
	fmt.Println("add2(1,2):", add2(1, 2))
	fmt.Println("add2(1.1,2.2):", add2(1.1, 2.2))
	fmt.Println("add2(\"Hello \",\"world\"):", add2("Hello ", "world"))
}
