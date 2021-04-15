package mymodule

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

func Loader2(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports2)
	// register other stuff
	L.SetField(mod, "name", lua.LString("my_module2"))
	L.SetField(mod, "field1", lua.LString("value1"))

	// returns the module
	L.Push(mod)
	return 1
}

var exports2 = map[string]lua.LGFunction{
	"myfunc": myfunc2,
}

func myfunc2(L *lua.LState) int {
	fmt.Println("call mymodule2 func[myfunc2()] in lua")
	return 0
}
