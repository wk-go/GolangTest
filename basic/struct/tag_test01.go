package main

/**
在处理json格式字符串的时候，经常会看到声明struct结构的时候，属性的右侧还有小米点括起来的内容。
在golang中，命名都是推荐都是用驼峰方式，并且在首字母大小写有特殊的语法含义：包外无法引用。但是由经常需要和其它的系统进行数据交互，例如转成json格式，存储到mongodb啊等等。这个时候如果用属性名来作为键值可能不一定会符合项目要求。

所以呢就多了小米点的内容，在golang中叫标签（Tag），在转换成其它数据格式的时候，会使用其中特定的字段作为键值。

那么当我们需要自己封装一些操作，需要用到Tag中的内容时，咋样去获取呢？这边可以使用反射包（reflect）中的方法来获取。
ps:http://www.01happy.com/golang-struct-tag-desc-and-get/
 */

import (
    "encoding/json"
    "fmt"
    "reflect"
)

func main() {
    type User struct {
        UserId   int    `json:"user_id" bson:"b_user_id"`
        UserName string `json:"user_name" bson:"b_user_name"`
    }
    // 输出json格式
    u := &User{UserId: 1, UserName: "tony"}
    j, _ := json.Marshal(u)
    fmt.Println(string(j))
    // 输出内容：{"user_id":1,"user_name":"tony"}

    // 获取tag中的内容 反射
    t := reflect.TypeOf(u)
    field := t.Elem().Field(0)
    fmt.Println(field.Tag.Get("json"))
    // 输出：user_id
    fmt.Println(field.Tag.Get("bson"))
    // 输出：user_id
}
