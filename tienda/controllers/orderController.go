package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection

func InitOrderController(client *mongo.Client) {
	orderCollection = client.Database("test").Collection("orders")
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var order bson.M
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Error procesando la orden", http.StatusBadRequest)
		return
	}

	order["date"] = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := orderCollection.InsertOne(ctx, order)
	if err != nil {
		http.Error(w, "Error registrando la orden", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Orden creada exitosamente"))
}
