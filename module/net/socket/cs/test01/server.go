package main

import(
	"fmt"
	"net"
	"io"
	"time"
)

func main(){
	ln, err := net.Listen("tcp", ":18081")
	if err != nil{
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	// Shut down the connection.
	defer conn.Close()

	fmt.Println("connect")
	var msg []byte
	b := make([]byte,8)
	if err := conn.SetReadDeadline(time.Now().Add(time.Millisecond*300)); err != nil {
		fmt.Println(err)
	}
	for {
		n, err := conn.Read(b)
		if err == io.EOF  || n == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println("n:", n)
		fmt.Println("b:",b[:len(b)])
		fmt.Println("b:",b[:n])
		fmt.Println("b string:",string(b[:n]))

		msg = append(msg,b[:n]...)
	}

	conn.Write(msg)
	// Echo all incoming data.
	//io.Copy(conn, conn)
}
