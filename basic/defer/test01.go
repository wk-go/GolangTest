package main
/*
https://www.oschina.net/news/107411/github-bullshit-code-with-go-defer
坑爹代码 | Go 语言的 defer 能制造出多少坑来？
 */
import "fmt"

type Slice []int

func NewSlice() Slice{
    return make(Slice, 0)
}

func(s *Slice) Add(elem int) *Slice{
    *s = append(*s, elem)
    fmt.Print(elem)
    return s
}

func main(){
    test1()
    fmt.Println("\n-------------")
    test2()
    fmt.Println("\n-------------")
    test3()
    fmt.Println("\n-------------")
    test4()
    fmt.Println("\n-------------")
    test5()
    fmt.Println("\n-------------")
}
func test1(){
    fmt.Println("test1:")
    s := NewSlice()
    defer s.Add(1)
    s.Add(3)
}
func test2(){
    fmt.Println("test2:")
    s := NewSlice()
    defer s.Add(1).Add(2)
    s.Add(3)
}
func test3(){
    fmt.Println("test3:")
    s := NewSlice()
    defer s.Add(1).Add(2).Add(3)
    s.Add(3)
}
func test4(){
    fmt.Println("test4:")
    s := NewSlice()
    defer s.Add(1).Add(2).Add(4)
    s.Add(3)
}
func test5(){
    fmt.Println("test5:")
    s := NewSlice()
    defer s.Add(1).Add(2).Add(4).Add(5)
    s.Add(3)
}