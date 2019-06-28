package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func main() {
	commandHandle()
}

func commandHandle() {
	test1App := kingpin.New("test1App", "test1App").Version("v1.0.1")
	Cmd1 := test1App.Command("Cmd1", "Cmd1")
	paramA := Cmd1.Flag("paramA", "").Default("1").String()
	Cmd2 := test1App.Command("Cmd2", "Cmd2")
	paramB := Cmd2.Flag("paramB", "").Default("2").String()

	Cmd2_1 := Cmd2.Command("Cmd2_1", "Cmd2_1")
	paramX1 := Cmd2_1.Flag("paramX1", "").Default("3").String()

	kingpin.MustParse(test1App.Parse(os.Args[1:]))
	fmt.Println("paraA:", *paramA)
	fmt.Println("paraB:", *paramB)
	fmt.Println("paraX1:", *paramX1)
}
