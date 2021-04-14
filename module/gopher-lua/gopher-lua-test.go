package main

import (
	"flag"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"os"
)

var (
	command    string
	filename   string
	commandMap = map[string]func(){
		"stringTest": stringTest,
	}
)

func main() {
	flag.StringVar(&command, "c", "stringTest", "command")
	flag.StringVar(&filename, "f", "", "file name")
	flag.Parse()
	flag.Usage = usage
	if command == "" {
		command = "stringTest"
	}
	if f, ok := commandMap[command]; ok {
		f()
	} else {
		fmt.Println("Command Not Found")
	}
}
func usage() {
	fmt.Fprintf(os.Stderr, `gopher-lua-test:
Usage: gopher-lua-test [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}

func stringTest() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`print("hello from lua")`); err != nil {
		panic(err)
	}
}
