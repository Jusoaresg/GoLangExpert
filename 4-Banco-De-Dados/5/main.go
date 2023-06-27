package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*

	O product HAS ONE from category
	O Category HAS MANY from Product

*/

type Category struct {
	ID       int
	Name     string
	Products []Product
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
	/*categorie := Category{Name: "Cozinha"}
	db.Create(&categorie)*/

	//Create Product
	db.Create(&Product{
		Name:       "Facas",
		Price:      50,
		CategoryID: 2,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name)
		}
	}
}
