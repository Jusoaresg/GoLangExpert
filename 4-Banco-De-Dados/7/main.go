package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*

	O product HAS MANY TO MANY from category
	da mesma forma que category HAS MANY TO MANY from products

*/

type Category struct {
	ID       int `gorm:"PrimaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories"`
}

type Product struct {
	ID         int `gorm:"PrimaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories"`
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
	categorie := Category{Name: "Cozinha"}
	db.Create(&categorie)

	categorie2 := Category{Name: "Eletronico"}
	db.Create(&categorie2)

	//Create Product
	db.Create(&Product{
		Name:       "Panela",
		Price:      99.0,
		Categories: []Category{categorie, categorie2},
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
