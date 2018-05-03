package main
// 定时器测试
import (
    "time"
    "fmt"
)

var input = make(chan string, 10)

func main(){
    demo()
}

func demo() {
    t1 := time.NewTimer(time.Second * 5)
    t2 := time.NewTimer(time.Second * 10)

    for {
        select {
        case msg := <- input:
            fmt.Println(msg)

        case <-t1.C:
            fmt.Println("5s timer")
            t1.Reset(time.Second * 5)

        case <-t2.C:
            fmt.Println("10s timer")
            t2.Reset(time.Second * 10)
        }
    }
}

