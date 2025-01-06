package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Product estructura para productos
type Product struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// Order estructura para pedidos
type Order struct {
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Products []string `json:"products"`
}

// GetAllProducts obtiene los productos desde el backend
func GetAllProducts() []Product {
	// Cambiamos a la ruta pública
	resp, err := http.Get("http://localhost:3000/admin/products")
	if err != nil {
		fmt.Println("Error al obtener productos:", err)
		return nil
	}
	defer resp.Body.Close()

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		fmt.Println("Error al decodificar productos:", err)
		return nil
	}
	return products
}


// CreateOrder envía un pedido al backend
func CreateOrder(name, address string, products []string) error {
	order := Order{Name: name, Address: address, Products: products}
	body, _ := json.Marshal(order)

	// URL corregida: POST a /admin/orders
	resp, err := http.Post("http://localhost:3000/admin/orders", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Error al crear el pedido: %s", resp.Status)
	}
	return nil
}

