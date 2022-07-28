package main

import (
	"fmt"
	"golang_test/basic/struct/test/pkg01"
)

//测试结构体的可导出属性
func main() {
	struct01 := new(pkg01.Struct01)
	struct01.Name = "zhangsan"
	struct01.Age = "20"
	//struct01._Message = "Hello world" // 不能访问
	//struct01.字段1 = "value1" // 不能访问
	struct01.F字段1 = "V值1"
	fmt.Printf("struct01:%#v\n", struct01)
	fmt.Printf("struct02:%#v\n", pkg01.NewStruct01("lisi", "22", "Hello world", "Hello word2", "值1", "V值1"))
}
