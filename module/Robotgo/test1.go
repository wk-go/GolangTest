package main

import (
	//. "fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.ScrollMouse(10, "up")
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
	robotgo.MouseClick("left",true)
}