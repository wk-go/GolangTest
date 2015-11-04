package main
import (
	"fmt"
	_ "time"
	"time"
)

var a int
var ch chan int

func f1(i int){
	fmt.Printf("Hello:%v\n",i)
	//time.Sleep(100*time.Millisecond)
	a=i
	<-ch
}


func main(){
	ch = make(chan int)
	for i := 0;i<20;i++{
		go f1(i)
	}

	for i :=0 ; i<20;i++{
		ch<-1
		fmt.Println(a)
	}

	c := time.Tick(1 *time.Second)
	for now := range c{
		fmt.Printf("%v\n",now)
	}
}