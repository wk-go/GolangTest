package main

var a string
var ch chan int

func f(){
	<-ch
	a = "Hello China!"
	print("Hello\n")
}

func main(){
	a = "Hello world!"
	ch = make(chan int)
	go f()
	ch<-0
	print(a)
}