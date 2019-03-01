package main

import (
    "gopkg.in/alecthomas/kingpin.v2"
    "fmt"
    "os"
)

func main(){
    app := kingpin.New("App name", "How to use")
    app.Author("WK")
    flag := app.Flag("flag", "A flag value").Short('f').Bool()
    name := app.Flag("name", "string name").Short('n').String()

    //sub command
    sub := app.Command("sub", "sub command")
    ip := sub.Flag("ip", "IP address").Short('i').String()
    port := sub.Flag("port", "Port").Short('p').String()

    //app.Parse(os.Args[1:])
    command := kingpin.MustParse(app.Parse(os.Args[1:]))
    fmt.Printf("Command: %+v\n", command)


    fmt.Printf("flag:%+v, name:%+v\n", *flag, *name)
    fmt.Printf("ip:%+v, port:%+v\n", *ip, *port)
}