// repository/beer_repository.go
package repository

import (
	"context"

	"github.com/d90ares/iBeers/internal/domain"
)

type Repository interface {
	GetAllBeers(ctx context.Context) ([]*domain.Beer, error)
	AddBeer(ctx context.Context, beer *domain.Beer) (*domain.Beer, error)
	GetByID(ctx context.Context, id int64) (*domain.Beer, error)
}
