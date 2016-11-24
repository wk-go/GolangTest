package main

import (
	. "fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	keve := robotgo.AddEvent("k")
	if keve == 0 {
		Println("you press...", "k")
	}

	mleft := robotgo.AddEvent("mleft")
	if mleft == 0 {
		Println("you press...", "mouse left button")
	}
}