package main

import(
	"github.com/robertkrimen/otto"
	"fmt"
)

func main(){
	// Create a VM
	vm := otto.New()

	// Run something in the VM
	vm.Run(`
		abc = 2+2;
		console.log("The value of abc is " + abc);
`)
	// Get a value out of the VM
	value, _ := vm.Get("abc")
	{
		value,_:= value.ToInteger()
		fmt.Println("value:", value)
	}

	// set a number
	vm.Set("def", 11)
	vm.Run(`
		console.log("The value of def is " + def);
		// The value of def is 11
`)

	// Set a string
	vm.Set("xyzzy", "Nothing happens.")
	vm.Run(`
		console.log(xyzzy.length); //16
`)

	// Get the value of the expression
	value,_ = vm.Run("xyzzy.length")
	{
		//value is an int64 with a value of 16
		value,_ := value.ToInteger()
		fmt.Println("value:",value)
	}

	// An error happens
	value, err := vm.Run("abcdefghijlmnopqrstuvwxyz.length")
	if err != nil{
		fmt.Println(err)
		fmt.Println(value.IsUndefined())
	}

	// Set a Go function
	vm.Set("sayHello", func(call otto.FunctionCall) otto.Value{
		fmt.Printf("Hello, %s.\n", call.Argument(0).String())
		return otto.Value{}
	})
	// Set a Go function that return something useful
	vm.Set("twoPlus", func(call otto.FunctionCall) otto.Value{
		right,_ := call.Argument(0).ToInteger()
		result,_ := vm.ToValue(2 + right)
		return result
	})
	// Use the functions in javascript
	result, _ := vm.Run(`
		sayHello("xyzzy");	// Hello, Xyzzy.
		sayHello(); 		// Hello, undefined
		
		result = twoPlus(2.0); //4
`)
	fmt.Println("result:", result)
}