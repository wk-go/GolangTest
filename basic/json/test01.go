package main

import (
    "time"
    "encoding/json"
    "fmt"
)
//通过设置标签，在转换成json字符串的时候，字段会转换成对应的小写
type Class struct {
    ClassName string    `json:"className"`
    Students []Person   `json:"students"`
}

type Person  struct{
    Name  string   `json:"name"`
    Age    int        `json:"age"`
    Time int64    `json:"-"`              // 直接忽略字段
}

func main(){
    class:=Class{
        ClassName:"三年一班",
        Students:[]Person{Person{"小明",18, time.Now().Unix()},Person{"李四",17, time.Now().Unix()},}}
    if result,err:=json.Marshal(&class);err==nil{
        fmt.Println("Marshal data:",string(result))


        classCopy := Class{}
        if err:= json.Unmarshal(result, &classCopy); err != nil{
            fmt.Println("Umarshal err:",err);
        }
        fmt.Printf("Umarshal data:%+v", classCopy)
    }

}