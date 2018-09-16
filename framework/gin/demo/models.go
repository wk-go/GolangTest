package main

import "github.com/jinzhu/gorm"

type Article struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    CatId   string `json:"cat_id"`
}

type Category struct {
    gorm.Model
    Name string `json:"name"`
}

type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"password"`
}
