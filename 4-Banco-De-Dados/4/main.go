package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*

	O produto HAS ONE from Serial Number

*/

type Category struct {
	ID   int
	Name string
}

type Product struct {
	ID           int `gorm:"PrimaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	//Create Categorie
	categorie := Category{Name: "Eletronicos"}
	db.Create(&categorie)

	//Create Product
	db.Create(&Product{
		Name:       "Keyboard",
		Price:      1000,
		CategoryID: 1,
	})

	//Create Serial Number
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category, product.SerialNumber.Number)
	}
}
