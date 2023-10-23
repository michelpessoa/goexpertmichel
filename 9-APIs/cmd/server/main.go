package main

import (
	"net/http"
	
	"github.com/michelpessoa/goexpertmichel/9-Apis/configs"
	"github.com/michelpessoa/goexpertmichel/9-Apis/internal/entity"
	"github.com/michelpessoa/goexpertmichel/9-Apis/internal/infra/database"
	"github.com/michelpessoa/goexpertmichel/9-Apis/internal/infra/webserver/handlers"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Forma encapsulada
	// config := configs.NewConf()

	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	
	http.ListenAndServe(":8000", r)

}

