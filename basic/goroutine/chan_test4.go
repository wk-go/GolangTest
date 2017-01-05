package main

// test the channel length status

import "fmt"

func main(){
    testChan := make(chan int, 3)
    fmt.Println("len(testChan):",len(testChan))
    testChan <- 1
    fmt.Println("len(testChan):",len(testChan))
    testChan <- 2
    fmt.Println("len(testChan):",len(testChan))
    testChan <- 3
    fmt.Println("len(testChan):",len(testChan))

    x := <- testChan
    fmt.Println("len(testChan):",len(testChan),"x:", x)
    x = <- testChan
    fmt.Println("len(testChan):",len(testChan),"x:", x)
    x = <- testChan
    fmt.Println("len(testChan):",len(testChan),"x:", x)
}
