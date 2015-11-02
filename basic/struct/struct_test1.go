package main
/*
测试
 */
import(
	"fmt"
	"reflect"
)

type test struct {
	in,out int
}

func main(){
	//value
	t1 := test{1,2}
	fmt.Println(reflect.TypeOf(t1), t1.in, t1.out)

	//pointer
	t2 := new(test)
	fmt.Println(reflect.TypeOf(t2), t2.in, t2.out)
}