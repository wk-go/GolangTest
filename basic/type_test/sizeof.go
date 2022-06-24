package main

/**
对齐保证:
对于任意类型的变量 x ，unsafe.Alignof(x) 至少为 1。
对于 struct 结构体类型的变量 x，计算 x 每一个字段 f 的 unsafe.Alignof(x.f)，unsafe.Alignof(x) 等于其中的最大值。
对于 array 数组类型的变量 x，unsafe.Alignof(x) 等于构成数组的元素类型的对齐倍数。
*/

import (
	"fmt"
	"unsafe"
)

type demo1 struct {
	n1 int8
	n2 int16
	n3 int32
	n4 int64
}

func main() {

	fmt.Println("unsafe.Sizeof(map[string]interface{}):", unsafe.Sizeof(map[string]interface{}{}))
	fmt.Println("unsafe.Sizeof(make(map[string]int)):", unsafe.Sizeof(make(map[string]int)))

	fmt.Println("unsafe.Sizeof(1.0) [float64]:", unsafe.Sizeof(1.0)) //float64
	fmt.Println("unsafe.Sizeof(float32(1.0)):", unsafe.Sizeof(float32(1.0)))
	fmt.Println("unsafe.Sizeof(byte(1)):", unsafe.Sizeof(byte(1)))
	fmt.Println("unsafe.Sizeof(rune(1)):", unsafe.Sizeof(rune(1)))
	fmt.Println("unsafe.Sizeof(1) [int]:", unsafe.Sizeof(1)) //int
	fmt.Println("unsafe.Sizeof(int8(1)):", unsafe.Sizeof(int8(1)))
	fmt.Println("unsafe.Sizeof(int16(1)):", unsafe.Sizeof(int16(1)))
	fmt.Println("unsafe.Sizeof(int32(1)):", unsafe.Sizeof(int32(1)))
	fmt.Println("unsafe.Sizeof(int64(1)):", unsafe.Sizeof(int64(1)))
	fmt.Println("unsafe.Sizeof(\"\"):", unsafe.Sizeof(""))
	fmt.Println("unsafe.Sizeof(\"hello world\"):", unsafe.Sizeof("hello world"))
	fmt.Println("unsafe.Sizeof([0]int{}):", unsafe.Sizeof([0]int{}))
	fmt.Println("unsafe.Sizeof(make([]int, 0)):", unsafe.Sizeof(make([]int, 0)))
	fmt.Println("unsafe.Sizeof(struct {}{}):", unsafe.Sizeof(struct{}{}))

	fmt.Println("unsafe.Sizeof(demo1{}):", unsafe.Sizeof(demo1{}))
}
