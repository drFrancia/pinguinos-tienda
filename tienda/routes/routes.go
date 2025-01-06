package routes

import (
	"net/http"

	"tienda/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/", controllers.GetProductsHandler)       // Ruta principal
	http.HandleFunc("/products", controllers.GetAllProductsHandler) // API para productos
	http.HandleFunc("/orders", controllers.CreateOrderHandler)      // API para órdenes
}
