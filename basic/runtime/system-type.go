package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("runtime.GOARCH:", runtime.GOARCH)
    fmt.Println("runtime.GOOS:", runtime.GOOS)
}
