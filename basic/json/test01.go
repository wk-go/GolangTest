package main

/**
<p>结构体必须是<strong>大写字母开头</strong>的成员才会被JSON处理到，小写字母开头的成员不会有影响。</p>
<p>Mashal时，结构体的成员变量名将会直接作为JSON Object的key打包成JSON；Unmashal时，会自动匹配对应的变量名进行赋，大小写不敏感。<br>
</p>
<p>Unmarshal时，如果JSON中有多余的字段，会被直接抛弃掉；如果JSON缺少某个字段，则直接忽略不对结构体中变量赋，不会报错。<br>
 */

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