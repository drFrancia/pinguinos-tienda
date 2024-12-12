package routes

import (
	"frontend/controllers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", controllers.ShowProducts)   // Página principal
	http.HandleFunc("/checkout", controllers.Checkout) // Página de pedidos
}
