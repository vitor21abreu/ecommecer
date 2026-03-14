package products

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Produto struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Qty         int                `bson:"qty" json:"qty"`
	Value       float64            `bson:"value" json:"value"`
	DueDate     time.Time          `bson:"due_date" json:"due_date"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}
