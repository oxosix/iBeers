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

func (s *BeerService) SearchBeer(ctx context.Context, string) (*domain.Beer, error) {
	beer, err := s.repository.Search()
	if err != nil {
		return nil, err
	}

	if beer == nil {
		return nil, errors.New("Beer not found")
	}
}

func (s *BeerService) StoreBeer(ctx context.Context, beer *domain.Beer) error {
	if beer, err := s.repository.Store(beer); err != nil {
		return nil, err
	}
	
}

func (s *BeerService) ValidateBeer(beer *domain.Beer) error {
	if beer.Name == "" {
		return errors.New("Name beer cannot be empty")
	}
	return nil

	if beer.Style || beer.Type == "" {
		return errors.New("Style and Type cannot be empty")
	}
	return nil

}
