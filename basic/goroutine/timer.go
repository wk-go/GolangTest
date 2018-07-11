package main

// 定时器测试
import (
	"time"
	"log"
)

var input = make(chan string, 10)

func main() {
	demo()
}

func demo() {
	t1 := time.NewTimer(time.Second * 6)
	t2 := time.NewTimer(time.Second * 10)
	t3 := time.NewTimer(time.Second * 2)

	for {
		select {
		case msg := <-input:
			log.Println(msg)

		case <-t1.C:
			log.Println("[06s] timer")
			t1.Reset(time.Second * 6)

		case <-t2.C:
			log.Println("[10s] timer")
			t2.Reset(time.Second * 10)

		case <-t3.C:
			input <- "[02s] timer:set input"
			t3.Reset(time.Second * 2)
		}

	}
}
