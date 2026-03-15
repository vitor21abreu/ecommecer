package catalog

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CatalogoRepository interface {
	Create(ctx context.Context, catalogo *Catalogo) (primitive.ObjectID, error)
	ListAll(Ctx context.Context) ([]Catalogo, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type catalogoRepositorio struct {
	collection *mongo.Collection
}

func NewCatalogoRepository(db *mongo.Database) CatalogoRepository {
	return &catalogoRepositorio{
		collection: db.Collection("catalogo"),
	}
}

func (r *catalogoRepositorio) Create(ctx context.Context, catalogo *Catalogo) (primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, catalogo)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil

}

func (r *catalogoRepositorio) ListAll(ctx context.Context) ([]Catalogo, error) {

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var catalogos []Catalogo
	if err := cursor.All(ctx, &catalogos); err != nil {
		return nil, err
	}

	return catalogos, nil

}

func (r *catalogoRepositorio) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
