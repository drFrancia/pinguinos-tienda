package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection

func InitProductController(client *mongo.Client) {
	productCollection = client.Database("test").Collection("products")
}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error obteniendo productos", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var products []bson.M
	if err := cursor.All(ctx, &products); err != nil {
		http.Error(w, "Error procesando productos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("views/index.html") // Ruta actualizada
    if err != nil {
        fmt.Println("Error al cargar la plantilla:", err)
        http.Error(w, "Error cargando la plantilla", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}
