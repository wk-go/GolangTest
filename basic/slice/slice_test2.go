package main

import "fmt"

/**
考点:
slice的底层数据结构是什么？给slice赋值，到底赋了什么内容？
通过:操作得到的新slice和原slice是什么关系？新slice的长度和容量是多少？
append在背后到底做了哪些事情？
slice的扩容机制是什么？
*/
func main() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	x = append(x, y...)
	x = append(x, y...)
	fmt.Println(a, x)
}
