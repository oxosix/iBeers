package domain

import "context"

type Repository interface {
	GetAllBeers(ctx context.Context) ([]*Beer, error)
	AddBeer(ctx context.Context, beer *Beer) (*Beer, error)
}

type UseCase interface {
	Repository
}
