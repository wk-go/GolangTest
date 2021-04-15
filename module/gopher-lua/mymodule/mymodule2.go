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
	"myfunc":  myfunc2,
	"sum_lua": sum2_lua,
}

func myfunc2(L *lua.LState) int {
	fmt.Println("call mymodule2 func[myfunc2()] in lua")
	return 0
}

func sum2(numbers ...int64) int64 {
	var sum int64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//多参数处理示例
func sum2_lua(L *lua.LState) int {
	numbers := make([]int64, 100, 100)
	i := 1
	for {
		lv := L.Get(i)
		if lv == lua.LNil {
			break
		}

		if v, ok := lv.(lua.LNumber); ok {
			numbers[i-1] = int64(v)
		}
		i++
	}
	sum := sum2(numbers[:i-1]...)
	L.Push(lua.LNumber(sum))
	return 1
}
