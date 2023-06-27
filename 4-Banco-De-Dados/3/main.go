package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int
	Name string
}

type Product struct {
	ID         int `gorm:"PrimaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	//Create Categorie
	/*categorie := Category{Name: "Eletronicos"}
	db.Create(&categorie)*/

	//Create Product
	/*db.Create(&Product{
		Name:       "Mouse",
		Price:      500,
		CategoryID: 1,
	})*/

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category)
	}
}
