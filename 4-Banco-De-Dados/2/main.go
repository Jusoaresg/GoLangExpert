package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	/*db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.00,
	})*/

	// Create Bash
	/*products := []Product{
		{Name: "Notebook", Price: 1000.00},
		{Name: "Mouse", Price: 50.00},
		{Name: "Keyboard", Price: 100.00},
	}
	db.Create(&products)*/

	// Select one
	/*var product Product
	db.First(&product, 17)
	db.First(&product, "name = ?", "Mouse")
	fmt.Println(product)*/

	// Select all
	/*var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}*/

	// Where
	//var products []Product
	//db.Where("Price > ?", 100).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	var products []Product
	db.Where("Name LIKE ?", "%book%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}