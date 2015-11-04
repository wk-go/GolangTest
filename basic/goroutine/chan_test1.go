// chan test 1
package main

import "fmt"

func sum(a []int, c chan int){
	sum :=0
	for _, v := range a {
		sum += v
	}
	c<-sum
}

func main(){
	a := []int{7,2,8,-9,4,0,1,3,4}

	c := make(chan int)
	fmt.Println(a[:len(a)/3])
	go sum(a[:len(a)/3], c)
	go sum(a[len(a)/3:len(a)/3*2],c)
	go sum(a[len(a)/3*2:],c)
	//x, y,z := <-c, <-c, <-c
	x := <-c
	y := <-c
	var z int
	if v, ok := <-c;ok{
		z = v
		fmt.Println(v)
	}

	fmt.Println(x,y,z,x+y+z)
}