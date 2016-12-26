package main

import(
    "github.com/bitly/go-simplejson"
    "fmt"
)


type profile struct {
    Nickname string
    Gender   int
}

func main()  {

    js,_ := simplejson.NewJson([]byte("{}"))
    js.Set("hello", "world")
    t,_ := js.MarshalJSON()
    fmt.Println("js.MarshalJSON():",string(t))
    m := map[string]interface{}{
        "name":"张三",
        "age":25,
        "friends":[]string{"赵四","杨素"},
        "profile":profile{"飞上天",24},
    }
    js.Set("userInfo",m)

    t,_ = js.MarshalJSON()
    fmt.Println("js.MarshalJSON():",string(t))

}