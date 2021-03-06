package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"golang_test/module/gopher-lua/data"
	"golang_test/module/gopher-lua/mymodule"
	luar "layeh.com/gopher-luar"
	"os"
)

var (
	command    string
	filename   string
	commandMap = map[string]func(){
		"stringTest":        stringTest,        // Do string
		"fileTest":          fileTest,          //Do file
		"callGoFromLua":     callGoFromLua,     // call lua function
		"callLuaFromGo":     callLuaFromGo,     // call lua function from go
		"callGoModule":      callGoModule,      // use module as lua module
		"useGoStruct":       useGoStruct,       // gopher-lua origin struct convert
		"useGoStructByLuar": useGoStructByLuar, //go struct convert to  lua table
		"gluamapperTest":    gluamapperTest,    //lua table convert to  go struct
		"codeToShareTest":   codeToShareTest,   //共享预编译脚本
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

func useGoStruct() {
	L := lua.NewState()
	defer L.Close()
	data.RegisterPersonType(L)
	if err := L.DoString(`
        p = person.new("Steeve")
        print(p:name()) -- "Steeve"
        p:name("Alice")
        print(p:name()) -- "Alice"
    `); err != nil {
		panic(err)
	}
}

func useGoStructByLuar() {
	L := lua.NewState()
	defer L.Close()

	script := `
print("Hello from Lua, " .. u.Name .. "!".."token:"..u:Token()) -- u.Token(u)
u:SetToken("12345")
print("Hello from Lua, " .. u.Name .. "!".."token changed:"..u:Token())
`

	u := &data.User{
		Name: "Tim",
	}
	u.SetToken("init_token")
	L.SetGlobal("u", luar.New(L, u))
	if err := L.DoString(script); err != nil {
		panic(err)
	}

	fmt.Println("Lua set your token to:", u.Token())
	// Output:
	// Hello from Lua, Tim!
	// Lua set your token to: 12345
}

// test module: github.com/yuin/gluamapper
func gluamapperTest() {
	type Role struct {
		Name string
	}

	type Person struct {
		Name      string
		Age       int
		WorkPlace string
		Role      []*Role
	}

	L := lua.NewState()
	if err := L.DoString(`
person = {
  name = "Michel",
  age  = "31", -- weakly input
  work_place = "San Jose",
  role = {
    {
      name = "Administrator"
    },
    {
      name = "Operator"
    }
  }
}
`); err != nil {
		panic(err)
	}
	var person Person
	if err := gluamapper.Map(L.GetGlobal("person").(*lua.LTable), &person); err != nil {
		panic(err)
	}
	fmt.Printf("Name:%s, Age:%d", person.Name, person.Age)
}

// CompileLua reads the passed lua file from disk and compiles it.
func CompileLua(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

// DoCompiledFile takes a FunctionProto, as returned by CompileLua, and runs it in the LState. It is equivalent
// to calling DoFile on the LState with the original source file.
func DoCompiledFile(L *lua.LState, proto *lua.FunctionProto) error {
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	return L.PCall(0, lua.MultRet, nil)
}

// Example shows how to share the compiled byte code from a lua script between multiple VMs.
func codeToShareTest() {
	codeToShare, _ := CompileLua("lua/codeToShare.lua")
	a := lua.NewState()
	b := lua.NewState()
	c := lua.NewState()
	DoCompiledFile(a, codeToShare)
	DoCompiledFile(b, codeToShare)
	DoCompiledFile(c, codeToShare)
}
