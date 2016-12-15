package main

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

type User struct {
    Id          int
    Name        string
    Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
    Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
    Id          int
    Age         int16
    User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
    Id    int
    Title string
    User  *User  `orm:"rel(fk)"`    //设置一对多关系
    Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
    Id    int
    Name  string
    Posts []*Post `orm:"reverse(many)"`
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(User), new(Profile), new(Tag), new(Post))
}

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    orm.RegisterDataBase("default", "mysql", "test:123@/2016-beego-test?charset=utf8")
}

func main() {
    o := orm.NewOrm()
    o.Using("default") // 默认使用 default，你可以指定为其他数据库

    profile := new(Profile)
    profile.Age = 30

    user := new(User)
    user.Profile = profile
    user.Name = "slene"

    fmt.Println(o.Insert(profile))
    fmt.Println(o.Insert(user))

    tag := new(Tag)
    tag.Name = "he"

    post := new(Post)
    post.Title="helle world"
    post.User = user
    post.Tags = append(post.Tags,tag)
    fmt.Println(o.Insert(post))
}