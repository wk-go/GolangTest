package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

type Article struct {
    Id   int
    Name string
}

func init() {
    orm.RegisterDriver("sqlite", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "data.db")
    orm.RegisterModel(new(Article))
}
func main() {

     //创建表  没成功，后来手动建表
    o := orm.NewOrm()
    art := Article{Name: "sea"}
    // 三个返回参数依次为：是否新创建的，对象Id值，错误
    if created, id, err := o.ReadOrCreate(&art, "Name"); err == nil {
        if created {
            fmt.Println("New Insert an object. Id:", id)
        } else {
            fmt.Println("Get an object. Id:", id)
        }
    }

    /*//写入数据 成功
    o := orm.NewOrm()
    art := new(Article)
    art.Name = "Mars"

    fmt.Println(o.Insert(art))
    */
/*
    // 查询数据 成功 不过发现id会自增 取出的是0，实际数据库中是1
    o := orm.NewOrm()
    art := Article{Name: "Mars"}
    err := o.Read(&art, "Name")

    if err == orm.ErrNoRows {
        fmt.Println("查询不到")
    } else if err == orm.ErrMissPK {
        fmt.Println("找不到主键")
    } else {
        fmt.Println(art.Id, art.Name)
    }*/

}