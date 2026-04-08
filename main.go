package main

import (
	"net/http"

	"github.com/FernandaIshida/go-web-application/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
