package main

import (
	"html/template"
	"log"
	"net/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// Producto estructura para los datos del backend
type Producto struct {
	ID          string  `json:"_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

// Renderizar plantillas
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("./views/" + tmpl)
	if err != nil {
		log.Println("Error cargando la plantilla:", err)
		http.Error(w, "Error cargando la plantilla", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func main() {
	// Ruta para archivos estáticos
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Página de inicio
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Obtener datos del backend
		resp, err := http.Get("http://localhost:3000/api/products")
		if err != nil {
			log.Println("Error al obtener productos:", err)
			http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Leer la respuesta JSON
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error al leer la respuesta:", err)
			http.Error(w, "Error al procesar productos", http.StatusInternalServerError)
			return
		}

		// Decodificar los datos
		var productos []Producto
		if err := json.Unmarshal(body, &productos); err != nil {
			log.Println("Error al decodificar productos:", err)
			http.Error(w, "Error al procesar productos", http.StatusInternalServerError)
			return
		}

		// Renderizar la plantilla
		renderTemplate(w, "index.pug", productos)
	})

	// Página de confirmación
	http.HandleFunc("/confirm", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Enviar pedido al backend
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("Error al leer el cuerpo del pedido:", err)
				http.Error(w, "Error al procesar pedido", http.StatusInternalServerError)
				return
			}

			resp, err := http.Post("http://localhost:3000/api/products", "application/json", bytes.NewReader(body))
			if err != nil {
				log.Println("Error al enviar pedido:", err)
				http.Error(w, "Error al confirmar pedido", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()
		}

		// Renderizar confirmación
		renderTemplate(w, "confirm.pug", nil)
	})

	// Iniciar servidor
	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
