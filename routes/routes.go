package routes

import (
	"net/http"

	"github.com/FernandaIshida/go-web-application/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
