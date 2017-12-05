package main

import (
    "fmt"
    "strconv"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println("local current time:",t)

    fmt.Println("utc time:", t.UTC().Format(time.UnixDate))

    fmt.Println("Unix timestamp:", t.Unix())

    timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
    fmt.Println("Utc Nano:",timestamp)
    timestamp = timestamp[:10]
    fmt.Println("Utc Unix timestamp:",timestamp)
}