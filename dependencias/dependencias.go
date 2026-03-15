package dependencias

import (
	"ecommecer/pkg/database"
	"ecommecer/internal/products"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

func Setup() *dig.Container {

	container := dig.New()

	container.Provide(database.NewMongoDataBase)

	container.Provide(func(db *mongo.Database) products.ProdutoRepositorio {
		return products.NovoProdutoRepositorio(db)
	})

	container.Provide(products.NovoProdutoServico)

	container.Provide(products.NovoProdutoHandler)

	return container
}
