package service

import (
	"context"
	"errors"

	"github.com/d90ares/iBeers/internal/ibeers/domain"
)

type BeerService struct {
	repository domain.Repository
}

func NewBeerService(repository domain.Repository) *BeerService {
	return &BeerService{
		repository: repository,
	}
}

func (s *BeerService) GetAllBeers(ctx context.Context) ([]*domain.Beer, error) {
	beers, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	if len(beers) == 0 {
		return nil, errors.New("No beers found")
	}

	return beers, nil
}

func (s *BeerService) GetBeerByID(ctx context.Context, ID int64) (*domain.Beer, error) {
	beer, err := s.repository.Get(ID)
	if err != nil {
		return nil, err
	}

	if beer == nil {
		return nil, errors.New("Beer not found")
	}

	return beer, nil
}
