package models

import "time"

type Order struct {
	ID       string    `bson:"_id,omitempty"`
	Name     string    `bson:"name"`
	Address  string    `bson:"address"`
	Products []string  `bson:"products"`
	Date     time.Time `bson:"date"`
}
