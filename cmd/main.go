package main

import (
	"ecommecer/dependencias"
	"ecommecer/internal/products"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	Start()

}

func Start() {

	container := dependencias.Setup()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	err := container.Invoke(func(produtoHandler *products.ProdutoHandler) error {

		router.POST("/produtos", produtoHandler.CriarProduto)
		router.GET("/produtos", produtoHandler.ListarProduto)
		router.PUT("/produtos/:id", produtoHandler.AlterarProduto)
		router.DELETE("/produtos/:id", produtoHandler.DeletarProduto)

		log.Println("Servidor rodando na porta :8080")

		return router.Run(":8080")
	})

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
