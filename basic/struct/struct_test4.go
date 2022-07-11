package main

// 属性与结构体重名测试效果

import "fmt"

type A struct {
	A string
}

type B struct {
	A
	B string
}

type C struct {
	B
}

func main() {
	c := &C{}
	c.A.A = "A value"
	c.B.B = "B value"
	fmt.Printf("A:%#v\n", c.A)
	fmt.Printf("B:%#v\n", c.B)
	fmt.Printf("C:%#v\n", c)
}
