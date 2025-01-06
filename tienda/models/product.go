package models

type Product struct {
	ID          string  `bson:"_id,omitempty"`
	Name        string  `bson:"name"`
	Price       float64 `bson:"price"`
	Description string  `bson:"description"`
}
