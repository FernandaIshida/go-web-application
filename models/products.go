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

	allProductsSelect, err := db.Query("SELECT *FROM products order by id asc")
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

		p.Id = id
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

func DeleteProduct(productId string) {
	db, err := db.DataBaseConnection()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(productId)
	defer db.Close()
}

func EditProduct(productId string) Product {
	db, err := db.DataBaseConnection()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	producInDataBase, err := db.Query("SELECT *FROM products WHERE id = $1", productId)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for producInDataBase.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = producInDataBase.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}
	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db, err := db.DataBaseConnection()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	updateProduct, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}
	updateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
