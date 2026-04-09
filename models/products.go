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

func CreateNewProduct(name, description string, price float64, quantity int) {
	db, err := db.DataBaseConnection()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	insertDataInBase, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDataInBase.Exec(name, description, price, quantity)
	defer db.Close()
}
