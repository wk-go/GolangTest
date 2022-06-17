package main

/**
	%v the value in a default format when printing structs, the plus flag (%+v) adds field names
	%#v a Go-syntax representation of the value
翻译过来就是：
	%v：会以默认形式打印值，当打印结构体的时候，"+"会打印字段名称
	%#v：值的Go语法表示法
*/

import "fmt"

type student struct {
	name string
	id   int
}

func main() {
	s := student{"Tom", 123}
	fmt.Printf("%%v的方式  : %v\n", s)
	fmt.Printf("%%+v的方式 : %+v\n", s)
	fmt.Printf("%%#v的方式 : %#v\n", s)
	fmt.Printf("%%v的方式(指针)  : %v\n", &s)
	fmt.Printf("%%+v的方式(指针) : %+v\n", &s)
	fmt.Printf("%%#v的方式(指针) : %#v\n", &s)
}
