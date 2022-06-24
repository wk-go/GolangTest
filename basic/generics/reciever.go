package main

//泛型 reciever
import "fmt"

type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

func main() {
	var intSlice MySlice[int] = []int{1, 2, 3, 4}
	fmt.Printf("intSlice.Sum():%d\n", intSlice.Sum())
	var float32Slice MySlice[float32] = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("float32Slice.Sum():%f\n", float32Slice.Sum())
}
