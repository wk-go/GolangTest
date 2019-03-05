package main

import "fmt"

/**
一趟快速排序的算法是：
1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
2）以第一个数组元素作为关键数据，赋值给key，即key=A[0]；
3）从j开始向前搜索，即由后开始向前搜索(j--)，找到第一个小于key的值A[j]，将A[j]的值赋给A[i]；
4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]的值赋给A[j]；
5）重复第3、4步，直到i=j； (3,4步中，没找到符合条件的值，即3中A[j]不小于key,4中A[i]不大于key的时候改变j、i的值，使得j=j-1，i=i+1，直至找到为止。找到符合条件的值，进行交换的时候i， j指针位置不变。另外，i==j这一过程一定正好是i+或j-完成的时候，此时令循环结束）。
 */

func QuickSort(array []int, left, right int){
    if left >= right{
        return
    }

    var (
        i   int = left
        j   int = right
        key int = array[left]
    )

    for i < j {
        for i < j && key <= array[j]{
            j--
        }
        array[i] = array[j]
        for i < j && key >= array[i]{
            i++
        }
        array[j] = array[i]
    }
    array[i] = key
    QuickSort(array, left, i - 1)
    QuickSort(array, i + 1, right)
}
/////////////////////////////////////////////目前这情况迭代不一定比递归快
type Stack struct {
    stack []int
    top   int
}
func NewStack() *Stack{
    return &Stack{
        stack: make([]int, 200),
        top:   0,
    }
}

func (s *Stack) Push(item1,item2 int){
    s.stack[s.top]= item1
    s.top += 1
    s.stack[s.top]= item2
    s.top += 1
}
func (s *Stack) Pop() (int,int){
    if s.Empty(){
        return -1, -1
    }
    s.top -= 2
    return s.stack[s.top], s.stack[s.top+1]
}

func (s *Stack) Empty() bool{
    return s.top == 0
}

func(s *Stack) Len() int{
    return s.top
}

func (s *Stack) Top() int{
    return s.top
}

func QuickSortIterator(array []int, left, right int){
    stack := NewStack()
    stack.Push(left,right)
    for !stack.Empty() {
        left,right = stack.Pop()
        if left >= right{
            continue
        }
        var (
            i   int = left
            j   int = right
            key int = array[left]
        )

        for i < j {
            for i < j && key <= array[j] {
                j--
            }
            array[i] = array[j]
            for i < j && key >= array[i] {
                i++
            }
            array[j] = array[i]
        }
        array[i] = key
        stack.Push(left, i-1)
        stack.Push(i+1, right)
    }
}

func main(){
    a := []int{10,23,4,5,6,35,33,66,9,7}
    fmt.Printf("%+v\n", a)
    QuickSort(a,0, len(a)-1)
    fmt.Printf("%+v\n", a)


    a = []int{10,23,4,5,6,35,33,66,9,7}
    fmt.Printf("%+v\n", a)
    QuickSortIterator(a,0, len(a)-1)
    fmt.Printf("%+v\n", a)
}