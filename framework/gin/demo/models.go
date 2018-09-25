package main

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Article struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time `json:"created_at" form:"-"`
    UpdatedAt time.Time `json:"updated_at" form:"-"`
    DeletedAt *time.Time `json:"deleted_at" sql:"index" form:"-"`
    Title   string `json:"title" form:"title" binding:"required"`
    Content string `json:"content" form:"content" binding:"required"`
    CatId   string `json:"cat_id" form:"cat_id" `
    Modified int64 `json:"modified" form:"-" gorm:"default:0"`
    Created int64 `json:"created" form:"-" gorm:"default:0"`
}
func (m *Article) BeforeSave()error{
    m.Modified = time.Now().Unix()
    if m.ID == 0{
        m.Created = m.Modified
    }
    return nil
}

type Category struct {
    gorm.Model
    Name string `json:"name" form:"name" binding:"required"`
}

type User struct {
    gorm.Model
    Username string `json:"username" gorm:"unique_index"`
    Password string `json:"password"`
}
