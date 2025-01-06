package main

import (
	"log"
	"net/http"
	"frontend/routes"
)

func main() {
	routes.RegisterRoutes()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
