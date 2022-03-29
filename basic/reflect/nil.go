package main

import (
	"fmt"
	"reflect"
)

// nil的反射问题

func main() {
	x := reflect.ValueOf((*struct{})(nil))
	fmt.Println("x:", x, x.IsNil())

	// 如果没有制定反射的值得类型就会报错
	y := reflect.ValueOf(nil)
	fmt.Println("y:", y)
	fmt.Println("y.IsNil", y.IsNil())
}
