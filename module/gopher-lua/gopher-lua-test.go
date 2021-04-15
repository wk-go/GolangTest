package main

import (
	"flag"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"golang_test/module/gopher-lua/mymodule"
	"os"
)

var (
	command    string
	filename   string
	commandMap = map[string]func(){
		"stringTest":    stringTest,
		"fileTest":      fileTest,
		"callGoFromLua": callGoFromLua,
		"callLuaFromGo": callLuaFromGo,
		"callGoModule":  callGoModule,
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
Usage: gopher-lua-test [-c command] [-f filename]

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

// run a lua file
func DoFile(L *lua.LState, f string) {
	if err := L.DoFile("lua/" + f); err != nil {
		panic(err)
	}
}

func fileTest() {
	L := lua.NewState()
	defer L.Close()
	if filename == "" {
		filename = "hello.lua"
	}
	DoFile(L, filename)
}

func Double(L *lua.LState) int {
	lv := L.ToInt(1)            /* get argument */
	L.Push(lua.LNumber(lv * 2)) /* push result */
	return 1                    /* number of results */
}

func callGoFromLua() {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("double", L.NewFunction(Double)) /* Original lua_setglobal uses stack... */
	if filename == "" {
		filename = "double.lua"
	}
	DoFile(L, filename)
}

func callLuaFromGo() {
	L := lua.NewState()
	defer L.Close()
	DoFile(L, "lua_double.lua")
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("lua_double"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(30)); err != nil {
		panic(err)
	}
	ret := L.Get(-1) // returned value
	L.Pop(1)
	fmt.Println("ret from lua:", ret, " ret.type():", ret.Type())
	if num, ok := ret.(lua.LNumber); ok {
		fmt.Printf("num: %T(%v)\n", num, num)
		num2 := float64(num)
		fmt.Printf("num2: %T(%v)\n", num2, num2)
	}
}

func callGoModule() {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("mymodule", mymodule.Loader)
	L.PreloadModule("mymodule2", mymodule.Loader2)
	DoFile(L, "call_go_module.lua")
}
