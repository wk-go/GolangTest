package main

import "fmt"

type Foo struct {
	Name string
}

// SetName 设置值到结构体要用指针
func (f *Foo) SetName(name string) {
	fmt.Printf("foo pointer in SetName: %p\n", f)
	f.Name = name
}

// GetName 获取值得时候可以使用值传递
func (f Foo) GetName() string {
	fmt.Printf("foo pointer in GetName: %p\n", &f)
	return f.Name
}

func (f Foo) String() string {
	return fmt.Sprintf("My name is %s", f.Name)
}

func main() {
	foo1 := &Foo{}
	fmt.Printf("foo pointer: %p\n", foo1)
	foo1.SetName("Sam")
	fmt.Println("foo1 name:", foo1.GetName())
	fmt.Println("Foo said", foo1)
}
