package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/FernandaIshida/go-web-application/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertPriceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price:", err)
		}

		convertQuantityToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity:", err)
		}

		models.CreateNewProduct(name, description, convertPriceToFloat, convertQuantityToInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error converting id to int:", err)
		}

		priceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price to float:", err)
		}

		quantityToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity to int:", err)
		}

		models.UpdateProduct(idToInt, name, description, priceToFloat, quantityToInt)
	}
	http.Redirect(w, r, "/", 301)
}
