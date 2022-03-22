package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 解法1
	go func() {
		t := time.NewTicker(time.Second * 1)
		for {
			// 1 在这里需要你写算法
			// 2 要求每秒钟调用一次proc函数
			// 3 要求程序不能退出
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@Solution1@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
							fmt.Println(err)
							var buf [4096]byte
							n := runtime.Stack(buf[:], false)
							fmt.Printf("==> %s\n", string(buf[:n]))
						}
					}()
					proc()
				}()
			}
		}
	}()

	//解法二
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		t := time.NewTicker(time.Second * 1)
		for {
			func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println("##########################Solution2########################")
						fmt.Println("", err)
					}
				}()
				proc()
			}()
			// time.Sleep(time.Second * 1)
			<-t.C
		}
	}()

	select {}
}

func proc() {
	panic("Boom!")
}
