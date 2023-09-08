package main

import (
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

	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.01,
	})

	db.Create(&[]Product{
		{Name: "Produto1", Price: 10.01},
		{Name: "Produto2", Price: 20.01},
		{Name: "Produto3", Price: 30.01},
		{Name: "Produto4", Price: 40.01},
	})

}
