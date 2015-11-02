/*
切片测试1
 */
package main

import(
	"fmt"
	"reflect"
)

func main(){
	mySlice1 := []int{0,1,2,3,4,5,6,7,8,9,10}
	fmt.Println(reflect.TypeOf(mySlice1),len(mySlice1),cap(mySlice1), mySlice1)

	//重新切片
	slice1 := mySlice1[:5]
	fmt.Println(reflect.TypeOf(slice1),len(slice1),cap(slice1),slice1)

	slice1 = mySlice1[2:5]
	fmt.Println(reflect.TypeOf(slice1),len(slice1),cap(slice1),slice1)

	//复制切片，以下两种方式结果相同
	slice1 = mySlice1[:len(mySlice1)]
	fmt.Println(reflect.TypeOf(slice1),len(slice1),cap(slice1),slice1)
	slice1 = mySlice1[:]
	fmt.Println(reflect.TypeOf(slice1),len(slice1),cap(slice1),slice1)

	//没有数据
	slice1 = mySlice1[1:1]
	fmt.Println(reflect.TypeOf(slice1),len(slice1),cap(slice1),slice1)


	//为各元素赋值
	slice2 := make([]int,10,20)
	fmt.Println(reflect.TypeOf(slice2),len(slice2),cap(slice2),slice2)
	//为各元素赋值,注意不要溢出 超出len()会报溢出错误
	for i := 0; i<10 ;i++{
		slice2[i] = i
	}
	fmt.Println(reflect.TypeOf(slice2),len(slice2),cap(slice2),slice2)
	//增加数据只能使用append
	for i := 100; i<120 ;i++{
		slice2 = append(slice2,i)
	}
	fmt.Println(reflect.TypeOf(slice2),len(slice2),cap(slice2),slice2)
}
