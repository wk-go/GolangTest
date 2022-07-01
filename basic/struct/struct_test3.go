package main

import "fmt"

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int
	n int
}

func main() {
	//匿名字段时基本数据类型的使用
	var e E
	e.Name = "Tom"
	e.Age = 300
	e.int = 20
	e.n = 40
	fmt.Println("e=", e)
	x := e.int
	fmt.Println("x=", x)
}
