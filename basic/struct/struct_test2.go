package main
/*
测试结构体 不设置字段名
 */
import(
	"fmt"
	"reflect"
)
//不必指定变量名
type test struct {
	int
	int8
	string
	people
}

type people struct{
	name string
	gender int
	address string
}

func main(){
	//value
	t1 := test{1,2,"2",people{name:"张三",address:"中国北京"}}
	fmt.Println(t1)
	//结构体重不设置字段名，使用类型名访问对应字段
	fmt.Println(reflect.TypeOf(t1), t1.int,t1.string)
	t1.people.name="李四"
	t1.people.gender=1
	fmt.Println(t1)

	//pointer
	t2 := new(test)
	fmt.Println(reflect.TypeOf(t2), t2)
}