package main

import (
	. "fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	//Println("test")
	abool := robotgo.ShowAlert("test", "robotgo")
	if abool == 0 {
		Println("ok@@@", "ok")
	}

	title:=robotgo.GetTitle()
	Println("title@@@", title)
}