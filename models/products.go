package models

import (
	"log"

	"github.com/FernandaIshida/go-web-application/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db, err := db.DataBaseConnection()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	allProductsSelect, err := db.Query("SELECT *FROM products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProductsSelect.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProductsSelect.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}
