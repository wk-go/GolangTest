package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 请使用两个goroutine交替打印1-100之间的奇数和偶数, 输出时按照从小到大输出。

func main() {
	fmt.Println("###### method 1 #####")
	method1()

	fmt.Println("###### method 2 #####")
	method2()

	fmt.Println("###### method 3 #####")
	method3()
}

func method1() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- struct{}{}
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
	wg.Wait()
}

// 用同步锁解决不了顺序问题
func method2() {
	runtime.GOMAXPROCS(1)
	locker := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			if i%2 == 1 {
				locker.Lock()
				fmt.Println(i)
				locker.Unlock()
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			if i%2 == 0 {
				locker.Lock()
				fmt.Println(i)
				locker.Unlock()
				runtime.Gosched()
			}
		}
	}()
	wg.Wait()
}

func method3() {
	runtime.GOMAXPROCS(1)
	ch1 := make(chan int)
	defer close(ch1)
	ch2 := make(chan int)
	defer close(ch2)

	run := func(ch chan int) {
		for {
			fmt.Println(<-ch)
			runtime.Gosched()
		}
	}
	go run(ch1)
	go run(ch2)

	for i := 1; i < 101; i++ {
		if i%2 == 1 {
			ch1 <- i
			continue
		}
		ch2 <- i
	}
}
