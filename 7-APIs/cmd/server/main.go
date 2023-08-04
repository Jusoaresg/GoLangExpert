package main

import (
	"github.com/Jusoaresg/GoLangExpert/7-APIs/configs"
	"github.com/Jusoaresg/GoLangExpert/7-APIs/internal/entity"
	"github.com/Jusoaresg/GoLangExpert/7-APIs/internal/infra/database"
	"github.com/Jusoaresg/GoLangExpert/7-APIs/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
