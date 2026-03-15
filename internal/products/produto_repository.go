package products

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProdutoRepositorio interface {
	Criar(ctx context.Context, produto *Produto) (primitive.ObjectID, error)
	Listar(ctx context.Context) ([]Produto, error)
	Alterar(ctx context.Context, id primitive.ObjectID, produto *Produto) error
	Deletar(ctx context.Context, id primitive.ObjectID) error
}

type produtoRepositorio struct {
	collection *mongo.Collection
}

func NovoProdutoRepositorio(db *mongo.Database) ProdutoRepositorio {
	return &produtoRepositorio{
		collection: db.Collection("produtos"),
	}
}

func (r *produtoRepositorio) Criar(ctx context.Context, produto *Produto) (primitive.ObjectID, error) {
	produto.CreatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, produto)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *produtoRepositorio) Listar(ctx context.Context) ([]Produto, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var produtos []Produto
	if err := cursor.All(ctx, &produtos); err != nil {
		return nil, err
	}

	return produtos, nil
}

func (r *produtoRepositorio) Alterar(ctx context.Context, id primitive.ObjectID, produto *Produto) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": produto})
	return err
}

func (r *produtoRepositorio) Deletar(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
