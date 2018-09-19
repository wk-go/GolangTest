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
