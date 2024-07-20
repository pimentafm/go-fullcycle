package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primary_key"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primary_key"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

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

	// var p Product
	// db.First(&p, 1)
	// p.Name = "Mobile Phone"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 2)
	// fmt.Println(p2)
	// db.Delete(&p2)

	// category := Category{Name: "Electronics"}
	// db.Create(&category)

	db.Create(&Product{
		Name:       "Notebook",
		Price:      20,
		CategoryID: 2,
	})

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	// var products []Product
	// db.Preload("Category").Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p.Name, p.Category.Name)
	// }

	var categories []Category
	err = db.Model(&Category{}).Preload("Category").Preload("Products").Find(&categories).Error
	if err != nil {
		fmt.Println(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println("-", p.Name, c.Name)
		}
	}
}
