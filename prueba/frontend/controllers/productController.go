package controllers

import (
	"frontend/models"
	"html/template"
	"log"
	"net/http"
)

// ShowProducts renderiza la lista de productos desde el backend
func ShowProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Cargando productos desde el backend")
	products := models.GetAllProducts()
	log.Print(products,"hola")

	if len(products) == 0 {
		http.Error(w, "No se pudieron cargar los productos. Inténtalo más tarde.", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("views/layout.html", "views/index.html")
	if err != nil {
		http.Error(w, "Error cargando las vistas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", products)
	if err != nil {
		http.Error(w, "Error renderizando la plantilla: "+err.Error(), http.StatusInternalServerError)
	}
}

// Checkout procesa el formulario para realizar un pedido
func Checkout(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts() // Obtener productos disponibles

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		address := r.FormValue("address")
		selectedProducts := r.Form["products"]

		err := models.CreateOrder(name, address, selectedProducts)
		if err != nil {
			http.Error(w, "Error creando el pedido: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Renderizar checkout.html con los productos disponibles
	tmpl, err := template.ParseFiles("views/layout.html", "views/checkout.html")
	if err != nil {
		http.Error(w, "Error cargando la vista: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "layout", products)
}
