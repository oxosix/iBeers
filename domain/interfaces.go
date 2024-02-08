package domain

import "context"

type Repository interface {
	GetAll() ([]*Beer, error)
	GetbyID(ctx context.Context) (*Beer, error)
	Add(ctx context.Context) (*Beer, error)
	// Adicionar outros métodos de repositório conforme necessário
}

type UseCase interface {
	GetAllBeers(ctx context.Context) ([]*Beer, error)
	AddBeer(ctx context.Context, beer *Beer) (*Beer, error)
	// Adicionar outros métodos de caso de uso conforme necessário
}
