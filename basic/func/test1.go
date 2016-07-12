package main
//函数类型可变函数测试
import (
	"fmt"
	"strconv"
)

func main(){
	var m = map[string]f{
		"t1":test1,
		"t2":test2,
		"t3":test3,
		"t4":func(b []byte){
			fmt.Println("test4:",string(b))
		},
	}
	s:=[]byte("hello world!!!")

	m["t1"](s)
	for i:=1;i<5;i++{
		//fmt.Println("t"+strconv.Itoa(i))
		m["t"+strconv.Itoa(i)](s)
	}

}

type f func([]byte)

func test1(b []byte){
	fmt.Println("test1:",string(b))
}
func test2(b []byte){
	fmt.Println("test2:",string(b))
}
func test3(b []byte){
	fmt.Println("test3:",string(b))
}