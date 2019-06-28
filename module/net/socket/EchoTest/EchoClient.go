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
var readMethod = flag.Int("read", 1, "Port")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		panic(conn)
	}
	fmt.Printf("Connectiong to %s:%s\n", *host, *port)
	var wg sync.WaitGroup
	wg.Add(2)

	go handleWrite(conn, &wg)
	if *readMethod == 1 {
		go handleRead(conn, &wg)
	} else {
		go handleRead2(conn, &wg)
	}
	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		_, err := conn.Write([]byte("hello " + strconv.Itoa(i) + "\n"))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for i := 0; i < 10; i++ {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(line)
	}
}

func handleRead2(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Wait()
	//reader := bufio.NewReader(conn)
	var b []byte = make([]byte, 1)
	var str string
	count := 0
	for {
		n, err := conn.Read(b)
		if err != nil {
			fmt.Println(err)
			break
		}
		if string(b[:n]) != "\n" {
			str += string(b[:n])
		} else {
			count++
			fmt.Printf("count_%02d:%s", count, str+string(b[:n]))
			str = ""
			if count == 10 {
				break
			}
		}
	}
}
