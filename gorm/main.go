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

	//var products []Product
	// db.Find(&product)
	// for _, p := range product {
	// 	fmt.Printf("Product: %+v\n", p)
	// }

	// db.Limit(2).Offset(2).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// db.Where("price > ?", 500).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// db.Where("name LIKE ?", "%o%").Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	var p Product
	db.First(&p, 1)
	p.Name = "Mobile Phone"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2)
	db.Delete(&p2)
}
