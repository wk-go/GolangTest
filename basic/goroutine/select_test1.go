package main

import (
	"fmt"
	"time"
)

func main() {
	var count, count_tick int = 1,1
	var d, d2, target time.Duration = 1000, 500, 10000
	tick := time.Tick(d * time.Millisecond)
	boom := time.After(target * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Printf("tick.%v\n",count_tick)
			count_tick ++
			count = 1
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Printf("%v   .\n", count)
			time.Sleep(d2 * time.Millisecond)
			count++
		}
	}
}