package models

type Product struct {
	Name        string
	Price       float64
	Description string
}

// Obtener productos de prueba
func GetAllProducts() []Product {
	return []Product{
		{Name: "Pescado", Price: 12.99, Description: "Fresco desde el Ártico"},
		{Name: "Hielo", Price: 5.50, Description: "Perfecto para mantener todo frío"},
		{Name: "Esmoquin de pingüino", Price: 50.00, Description: "¡El último grito de la moda!"},
	}
}
