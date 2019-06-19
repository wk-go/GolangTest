package main

import (
    "flag"
    "fmt"
    "io"
    "net"
)

var host = flag.String("host", "", "IP address")
var port = flag.String("port", "8000", "Port")

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
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn){
    defer conn.Close()
    for{
        io.Copy(conn, conn)
    }
}
