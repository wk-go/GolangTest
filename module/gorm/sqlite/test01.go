package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

type Product struct{
	gorm.Model
	Code string
	Price uint
}
func main(){
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price:1000})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	fmt.Println(product)

	var product1 Product
	db.First(&product1, "code = ?", "L1212") // find product with code L1212
	fmt.Println(product1)

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	//got list
	var products []Product
	db.Find(&products)
	for i, model := range products{
		fmt.Printf("products[%d]:%v\n", i, model)
	}
	fmt.Println("Use Pointer")
	var productsPtr []*Product
	db.Find(&productsPtr)
	for i, model := range productsPtr{
		fmt.Printf("products[%d]:%v\n", i, model)
	}

	// Delete - delete product
	//db.Delete(&product)
}
