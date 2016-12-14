package main

import (
    "container/list"
    "fmt"
)

type Class struct {
    Id    string
    Name    string
    Std     *list.List
}
type Student struct  {
    Name    string
    gender  string
    Class   string
}
func main() {
    notPointer()
    fmt.Println("===============华丽的分割线==============")
    usePointer()
}

func notPointer()  {
    classes := list.New()
    classes.PushBack(Class{"class-1","三年一班",list.New()})
    classes.PushBack(Class{"class-2","四年二班",list.New()})
    classes.PushBack(Class{"class-3","五年三班",list.New()})
    for sub := classes.Front();sub!=nil;sub=sub.Next(){
        insertData(sub)
    }
    for sub := classes.Front();sub!=nil;sub=sub.Next(){
        for s_sub := sub.Value.(Class).Std.Front();s_sub != nil;s_sub = s_sub.Next(){
            fmt.Println(s_sub.Value)
        }
        fmt.Println("\n\n");
    }
}

func usePointer()  {
    classes := list.New()
    classes.PushBack(&Class{"class-1","三年一班",list.New()})
    classes.PushBack(&Class{"class-2","四年二班",list.New()})
    classes.PushBack(&Class{"class-3","五年三班",list.New()})
    for sub := classes.Front();sub!=nil;sub=sub.Next(){
        sub.Value.(*Class).Std.PushBack(&Student{Name:sub.Value.(*Class).Name+" 张三",gender:"男", Class:sub.Value.(*Class).Name})
        sub.Value.(*Class).Std.PushBack(&Student{Name:sub.Value.(*Class).Name+" 王昭君",gender:"女", Class:sub.Value.(*Class).Name})
        sub.Value.(*Class).Std.PushBack(&Student{Name:sub.Value.(*Class).Name+" 杨光",gender:"男", Class:sub.Value.(*Class).Name})
    }
    for sub := classes.Front();sub!=nil;sub=sub.Next(){
        for s_sub := sub.Value.(*Class).Std.Front();s_sub != nil;s_sub = s_sub.Next(){
            fmt.Println(s_sub.Value.(*Student))
        }
        fmt.Println("\n\n");
    }
}

func insertData(e *list.Element){
    e.Value.(Class).Std.PushBack(Student{Name:e.Value.(Class).Name+" 张三",gender:"男", Class:e.Value.(Class).Name})
    e.Value.(Class).Std.PushBack(Student{Name:e.Value.(Class).Name+" 王昭君",gender:"女", Class:e.Value.(Class).Name})
    e.Value.(Class).Std.PushBack(Student{Name:e.Value.(Class).Name+" 杨光",gender:"男", Class:e.Value.(Class).Name})
}