package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func main() {
	fmt.Println("Hello world")
	argvHandle()
}

func argvHandle() {
	test1App := kingpin.New("test1App", "test1App")
	Cmd1 := test1App.Command("Cmd1", "Cmd1")
	paramA := Cmd1.Flag("paramA", "").Default("1").String()
	Cmd2 := test1App.Command("Cmd2", "Cmd2")
	paramB := Cmd2.Flag("paramA", "").Default("2").String()

	test2App := kingpin.New("test2App", "test2App")
	test2Cmd := test2App.Command("test2Cmd", "test2Cmd")
	paramX := test2Cmd.Flag("paramA", "").Default("3").String()

	//如果两个app共存只能这种方式处理？
	if os.Args[1] == "test1App" {
		kingpin.MustParse(test1App.Parse(os.Args[2:]))
		fmt.Println("paraA:", *paramA)
		fmt.Println("paraB:", *paramB)
	} else {
		kingpin.MustParse(test2App.Parse(os.Args[2:]))
		fmt.Println("paraX:", *paramX)
	}
}
