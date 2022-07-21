package main

import "fmt"

// key不存在是返回值类型的零值
type test struct {
	Name string
}

func main() {
	mapTest := map[string]bool{
		"A": true,
		"B": true,
	}
	fmt.Printf("A:%#v\n", mapTest["A"])
	fmt.Printf("B:%#v\n", mapTest["B"])
	fmt.Printf("C:%#v\n", mapTest["C"])

	mapTest2 := map[string]test{
		"A": {"A"},
		"B": {"B"},
	}
	fmt.Printf("A:%#v\n", mapTest2["A"])
	fmt.Printf("B:%#v\n", mapTest2["B"])
	fmt.Printf("C:%#v\n", mapTest2["C"])

}
