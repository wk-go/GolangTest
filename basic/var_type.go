package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello world")
	var v1 int32 = 1
	var v2 int8 = 125
	fmt.Println(v1)
	fmt.Println(v2)

	v3, v4 := 258, 65536
	fmt.Println(v3)
	fmt.Println(v4)

	if v3 == v4 {
		fmt.Println("v4 与 v3 相等")
	} else {
		fmt.Println("v4 与 v3 不相等")
	}

	//浮点数
	var f1 float32 = 100
	f2 := 123456.654321 // float64
	fmt.Println(f1)
	fmt.Println(f2)

	//复数
	var cmpx1 complex64 = 12 + 3i
	cmpx2 := 16 + 5i
	fmt.Println(cmpx1)
	fmt.Println(cmpx2)

	//字符串
	str1 := "Hello world!"
	ch1 := str1[0]
	fmt.Println(str1)
	fmt.Println(ch1)
	fmt.Printf("The length of \"%s\" is %d \n", str1, len(str1))
	fmt.Printf("The first character of \"%s\" is %c.\n", str1, ch1)

	str := "Hello,世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}

	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune
	}

	//数组
	// 先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println()

	fmt.Println("\nElements of mySlice: ")
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println()

	mySlice = myArray[2:8]
	fmt.Println("\nElements of mySlice new: ")
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println()

	mySlice = make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))

	//字符串和数值转换
	//整型转换成字符串
	str = strconv.Itoa(100)
	fmt.Printf("type:%T value:%#v\n", str, str)

	//字符串转整型
	str1 = "110"
	str2 := "s100"
	str3 := "100abc"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("%v 转换失败！\n", str1)
	} else {
		fmt.Printf("type:%T value:%#v\n", num1, num1)
	}
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("%v 转换失败！\n", str2)
	} else {
		fmt.Printf("type:%T value:%#v\n", num2, num2)
	}

	num3, err := strconv.Atoi(str3)
	if err != nil {
		fmt.Printf("%v 转换失败！\n", str3)
	} else {
		fmt.Printf("type:%T value:%#v\n", num3, num3)
	}

	fstr1 := "1.2"
	fnum1, err := strconv.ParseFloat(fstr1, 64)
	if err != nil {
		fmt.Printf("%v 转换失败！\n", fstr1)
	} else {
		fmt.Printf("type:%T value:%#v\n", fnum1, fnum1)
	}

}
