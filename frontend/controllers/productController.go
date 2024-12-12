package controllers

import (
	"frontend/models"
	"html/template"
	"log"
	"net/http"
)

// Mostrar productos
func ShowProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("ShowProducts ejecutado")
	products := models.GetAllProducts()

	// Verifica si las plantillas existen y se cargan correctamente
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

// Procesar pedidos
func Checkout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Procesa el pedido y redirige a la página principal
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("views/layout.html", "views/checkout.html")
	if err != nil {
		http.Error(w, "Error cargando las vistas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", nil)
}
