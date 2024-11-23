// service/beer_service.go
package service

import (
	"context"

	"github.com/d90ares/iBeers/internal/app/repository"
	"github.com/d90ares/iBeers/internal/domain"
)

type BeerService struct {
	repository *repository.BeerRepository
}

func NewBeerService(repository *repository.BeerRepository) *BeerService {
	return &BeerService{
		repository: repository,
	}
}

func (s *BeerService) GetAllBeers(ctx context.Context) ([]*domain.Beer, error) {
	// Chame o método correspondente no repositório para obter todas as cervejas
	beers, err := s.repository.GetAll(ctx)
	if err != nil {
		// Se ocorrer um erro ao acessar o repositório, você pode lidar com isso aqui
		return nil, err
	}

	// Aqui você pode realizar qualquer lógica adicional se necessário

	return beers, nil
}

// Implementar outros métodos de serviço, como GetBeerByID, SearchBeer, StoreBeer, etc.
func (s *BeerService) AddBeer(ctx context.Context, beer *domain.Beer) (*domain.Beer, error) {
	beers, err := s.repository.Add(ctx, beer)
	if err != nil {
		return nil, err
	}
	return beers, nil
}

func (s *BeerService) GetByID(ctx context.Context, beer int64) (*domain.Beer, error) {
	beers, err := s.repository.GetByID(ctx, beer)
	if err != nil {
		return nil, err
	}
	return beers, nil
}
