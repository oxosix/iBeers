// usecase/beer_usecase.go
package usecase

import (
	"context"

	"github.com/d90ares/iBeers/internal/ibeers/domain"
	"github.com/d90ares/iBeers/internal/ibeers/service"
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

// Implementar outros métodos de caso de uso, como GetBeerByID, SearchBeer, StoreBeer, etc.
