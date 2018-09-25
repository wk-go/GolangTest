package main

import (
    "time"
)
type MyModel struct{
    ID        uint `gorm:"primary_key"`
    Modified int64 `json:"modified" form:"-" gorm:"default:0"`
    Created int64 `json:"created" form:"-" gorm:"default:0"`
}
func (m *MyModel) BeforeSave()error{
    m.Modified = time.Now().Unix()
    if m.ID == 0{
        m.Created = m.Modified
    }
    return nil
}

type Article struct {
    MyModel
    Title   string `json:"title" form:"title" binding:"required"`
    Content string `json:"content" form:"content" binding:"required"`
    CatId   string `json:"cat_id" form:"cat_id" `
}


type Category struct {
    MyModel
    Name string `json:"name" form:"name" binding:"required"`
}

type User struct {
    MyModel
    Username string `json:"username" gorm:"unique_index"`
    Password string `json:"password"`
}
