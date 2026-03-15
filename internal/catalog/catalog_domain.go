package catalog

import (
	"ecommecer/internal/products"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Catalog struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name" json:"name"`
	Products []primitive.ObjectID `bson:"products" json:"products"`
	itens    []products.Produto   `bson:"-" json:"-"`
}
