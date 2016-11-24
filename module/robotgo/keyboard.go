package main

import (
	//. "fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.TypeString("Hello World")
	robotgo.KeyTap("enter")
	robotgo.TypeString("en")
	robotgo.KeyTap("i", "alt", "command")
	arr := []string{"alt", "command"}
	robotgo.KeyTap("i", arr)
}