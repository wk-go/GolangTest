package main

import (
    "bufio"
    "flag"
    "fmt"
    "net"
    "strconv"
    "sync"
)

var host = flag.String("host", "localhost", "Host")
var port = flag.String("port", "8000", "Port")

func main(){
    flag.Parse()
    conn, err := net.Dial("tcp", *host + ":" + *port)
    if err != nil {
        panic(conn)
    }
    fmt.Printf("Connectiong to %s:%s\n", *host, *port)
    var wg sync.WaitGroup
    wg.Add(2)

    go handleWrite(conn, &wg)
    go handleRead(conn, &wg)
    wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup){
    defer wg.Wait()
    for i:=0; i< 10; i++{
        _, err :=conn.Write([]byte("hello " + strconv.Itoa(i) +"\r\n"))
        if err != nil {
            fmt.Println(err)
            break
        }
    }
}
func handleRead(conn net.Conn, wg *sync.WaitGroup){
    defer wg.Wait()
    reader := bufio.NewReader(conn)
    for i:=0; i < 10; i++{
        line, err := reader.ReadString(byte('\n'))
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(line)
    }
}