package main

import "fmt"

/**
冒泡排序算法的原理如下：
    1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
    2. 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。
    3. 针对所有的元素重复以上的步骤，除了最后一个。
    4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
 */

func BubbleSort(a []int){
    flag := true
    for j:=len(a)-1;j > 0; j--{
        flag = true
        for i:=0; i<j; i++{
            if a[i] > a[i+1]{
                a[i], a[i+1] = a[i+1], a[i]
                flag = false
                continue
            }
        }
        if flag {
            continue
        }
    }
}

func main(){
    a := []int{10,3,2,3,5,5,85,45,67, 80, 90, 200, 70}
    fmt.Printf("%+v\n", a)
    BubbleSort(a)
    fmt.Printf("%+v\n", a)
}
