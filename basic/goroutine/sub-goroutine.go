package main

import (
	"log"
	"time"
)

var block = make(chan int)

func main() {
	defer func() {
		log.Println("main exited")
	}()
	go routine1()
	<-block
}

func routine1() {
	defer func() {
		log.Println("routine1 exited")
	}()
	go routine2()
}

func routine2() {
	defer func() {
		log.Println("routine2 exited")
	}()
	count := 0
	for {
		if count == 5 {
			block <- 1
		}
		log.Println("running!!")
		time.Sleep(2 * time.Second)
		count++
	}
}
