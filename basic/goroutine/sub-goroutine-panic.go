package main

import (
	"log"
	"time"
)

func main() {
	go routine01()
	select {}
}

func routine01() {
	defer func() {
		log.Println("routine01 exited")
	}()
	go routine02()
}

func routine02() {
	defer func() {
		log.Println("routine02 exited")
	}()
	count := 0
	for {
		if count == 5 {
			panic("exited")
		}
		log.Println("running!!")
		time.Sleep(2 * time.Second)
		count++
	}
}
