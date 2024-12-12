package main

import (
	"log"
	"net/http"

	"frontend/routes"
)

func main() {
	// Registrar rutas
	routes.RegisterRoutes()

	// Servir archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Iniciar el servidor
	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
