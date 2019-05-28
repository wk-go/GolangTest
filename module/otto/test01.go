package main

import (
    "github.com/robertkrimen/otto"
    "fmt"
)

func main(){
    vm := otto.New()
    // Run something in the VM
    fmt.Println("Run somthing in the VM")
    vm.Run(`
        abc = 2 + 2;
        console.log("The value of abc is " + abc);//4
    `)

    // Get a value out of VM
    fmt.Println("Get a value out of VM")
    if value, err := vm.Get("abc"); err == nil {
        if valueInt, err := value.ToInteger(); err == nil{
            fmt.Printf("%#v %#v\n", valueInt, err)
        }
    }

    // Set a number
    vm.Set("def",11)
    vm.Run(`
        console.log("The Value of def is " + def);
    `)

    // Set a string
    vm.Set("xyzzy", "Nothing happens.")
    vm.Run(`
        console.log("The length of xyzzy is " + xyzzy.length);//16
    `)
    // Get the value of an expression
    fmt.Println("Get the value of an expression")
    value, _ := vm.Run("xyzzy.length")
    {
        valueInt, _ := value.ToInteger()
        fmt.Printf("The length of zyzzy is %d \n", valueInt)
    }
    // Set a Go function
    vm.Set("sayHello", func(call otto.FunctionCall) otto.Value{
        fmt.Printf("Hello, %s.\n", call.Argument(0).String())
        return otto.Value{}
    })
    // Set a Go function that returns something useful
    vm.Set("twoPlus", func(call otto.FunctionCall)otto.Value{
        right, _ := call.Argument(0).ToInteger()
        result, _ := vm.ToValue(2 + right)
        return result
    })
    // Use the functions in JavaScript
    result, _ := vm.Run(`
        sayHello("Xyzzy"); // Hello, Xyzzy.
        sayHello(); // Hello, undefined.
        result = twoPlus(2.0); //4
    `)
    resultInt, _ := result.ToInteger()
    fmt.Printf("The value of result is %d\n", resultInt)
}