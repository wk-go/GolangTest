package main

// golang中只有值传递

import (
	"log"
)

func main() {
	var _slice = []int{1, 2, 3, 5, 5}
	log.Printf("_slice value:%v\n",
		_slice)
	log.Printf("_slice pointer:%p\n", &_slice)
	update(_slice)
	log.Printf("_slice value:%v\n", _slice)

	log.Printf("\n\nTest Int\n\n")

	var _int = 1
	log.Printf("_int value:%v", _int)
	log.Printf("_int pointer:%p", &_int)
	add(&_int)
	log.Printf("_int value:%v", _int)
}

func update(a []int) {
	a[1] += 1
	log.Printf("a pointer:%p\n", &a)
}

func add(i *int) {
	*i = 1
	// 地址作为值传递给i
	log.Printf("i value:%v\n", i)
	log.Printf("i pointer:%p\n", &i)
}
