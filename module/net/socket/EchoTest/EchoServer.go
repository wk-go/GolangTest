package main

import (
    "bytes"
    "flag"
    "fmt"
    "io"
    "net"
)

var host = flag.String("host", "", "IP address")
var port = flag.String("port", "8000", "Port")
var handleType = flag.Int("handle", 1, "Port")

func main(){
    flag.Parse()
    l, err := net.Listen("tcp", *host + ":" + *port)
    if err != nil {
        panic(err)
    }
    defer l.Close()
    fmt.Printf("Listen on %s:%s\n", *host, *port)
    for {
        conn, err := l.Accept()
        if err != nil{
            fmt.Println(err)
        }
        if *handleType == 1 {
            go handleRequest(conn)
        }else{
            go handleRequest2(conn)
        }
    }
}

func handleRequest(conn net.Conn){
    defer conn.Close()
    for{
        io.Copy(conn, conn)
    }
}

func handleRequest2(conn net.Conn){
    defer conn.Close()
    buffSize := 3
    b :=make([]byte, buffSize)
    count := 0
    var str []byte
    for{
        n, err := conn.Read(b)
        if err != nil {
            fmt.Println(err)
        }
        //conn.Write(b[:n])

        index := bytes.Index(b,[]byte("\n"))
        if index !=-1 {
            str = append(str, b[:index+1]...)
        }else{
            str = append(str, b[:n]...)
        }
        if bytes.Index(str, []byte("\n")) != -1{
            conn.Write(str)
            count ++
            fmt.Printf("count_%02d:%v,%s",count, str,str)
            str=make([]byte,0,buffSize)
        }
        if index != -1 && index+1 < n{
            str=append(str,b[index+1:n]...)
        }
    }
}
