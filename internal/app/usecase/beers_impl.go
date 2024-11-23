package usecase

import (
	"context"

	"github.com/d90ares/iBeers/internal/app/service"
	"github.com/d90ares/iBeers/internal/domain"
)

type BeerUseCase struct {
	service *service.BeerService
}

func NewBeerUseCase(service *service.BeerService) *BeerUseCase {
	return &BeerUseCase{
		service: service,
	}
}

func (uc *BeerUseCase) GetAllBeers(ctx context.Context) ([]*domain.Beer, error) {
	// Implementar lógica para obter todas as cervejas usando o repositório
	beers, err := uc.service.GetAllBeers(ctx)
	if err != nil {
		return nil, err
	}

	return beers, nil
}

func (uc *BeerUseCase) AddBeer(ctx context.Context, beer *domain.Beer) (*domain.Beer, error) {
	beer, err := uc.service.AddBeer(ctx, beer)
	if err != nil {
		return nil, err
	}
	return beer, nil
}

func (uc *BeerUseCase) GetByID(ctx context.Context, id int64) (*domain.Beer, error) {
	beer, err := uc.service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return beer, nil
}

// Implementar outros métodos de caso de uso, como GetBeerByID, SearchBeer, StoreBeer, etc.
