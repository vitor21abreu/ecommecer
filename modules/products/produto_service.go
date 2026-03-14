package products

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProdutoService struct {
	repo ProdutoRepositorio
}

func NovoProdutoServico(repo ProdutoRepositorio) *ProdutoService {
	return &ProdutoService{repo: repo}
}

func (s *ProdutoService) CriarProduto(ctx context.Context, produto *Produto) (primitive.ObjectID, error) {
	return s.repo.Criar(ctx, produto)
}

func (s *ProdutoService) ListarProduto(ctx context.Context) ([]Produto, error) {
	return s.repo.Listar(ctx)
}

func (s *ProdutoService) AlterarProduto(ctx context.Context, id primitive.ObjectID, produto *Produto) error {
	return s.repo.Alterar(ctx, id, produto)
}

func (s *ProdutoService) DeletarProduto(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Deletar(ctx, id)
}
