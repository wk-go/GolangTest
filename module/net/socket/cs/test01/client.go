package main

import (
	"net"
	"fmt"
	"bufio"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:18081")

	if err != nil{
		panic(err)
	}
	fmt.Println("send")
	fmt.Fprintf(conn, "hello world")
	fmt.Println("send end")
	var b = make([]byte,1024)
	n, err := bufio.NewReader(conn).Read(b)
	fmt.Println("num:", n)
	if err != nil {
		panic(err)
	}
	fmt.Println("msg:",string(b))
}
