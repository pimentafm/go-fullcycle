package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	// products := []Product{
	// 	{Name: "Mobile", Price: 500},
	// 	{Name: "Tablet", Price: 300},
	// 	{Name: "Desktop", Price: 1500},
	// 	{Name: "Smartwatch", Price: 200},
	// }

	// db.Create(&products)

	//var product Product
	// db.First(&product, 1)
	// fmt.Printf("Product: %+v\n", product)

	// db.First(&product, "name = ?", "Tablet")
	// fmt.Printf("Product: %+v\n", product)

	var product []Product
	db.Find(&product)
	for _, p := range product {
		fmt.Printf("Product: %+v\n", p)
	}
}
