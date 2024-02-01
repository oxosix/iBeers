package usecase

import (
	"context"

	"github.com/d90ares/iBeers/internal/ibeers/domain"
	"github.com/d90ares/iBeers/internal/ibeers/service"
)

type BeerUseCase struct {
	service service.BeerService
}

func NewBeerUseCase(service service.BeerService) *BeerUseCase {
	return &BeerUseCase{
		service: service,
	}
}

func (uc *BeerUseCase) GetAllBeers(ctx context.Context) ([]*domain.Beer, error) {
	return uc.service.GetAllBeers(ctx)
}

func (uc *BeerUseCase) GetBeerByID(ctx context.Context, ID int64) (*domain.Beer, error) {
	return uc.service.GetBeerByID(ctx, ID)
}

func (uc *BeerUseCase) SearchBeer(ctx context.Context, b string) (*domain.Beer, error) {
	return uc.service.SearchBeer(ctx, b)
}

func (uc *BeerUseCase) StoreBeer(ctx context.Context, beer *domain.Beer) error {
	return uc.service.StoreBeer(ctx, beer)
}

func (uc *BeerUseCase) ValidateBeer(beer *domain.Beer) error {
	return uc.service.ValidateBeer(beer)
}
