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

func (s *BeerService) SearchBeer(ctx context.Context, b string) (*domain.Beer, error) {
	beer, err := s.repository.Search(b)
	if err != nil {
		return nil, err
	}

	if beer == nil {
		return nil, errors.New("Beer not found")
	}
	return beer, nil
}

func (s *BeerService) StoreBeer(ctx context.Context, beer *domain.Beer) error {
	if err := s.repository.Store(beer); err != nil {
		return err
	}
	return nil
}

func (s *BeerService) ValidateBeer(beer *domain.Beer) error {
	if beer.Name == "" {
		return errors.New("Name beer cannot be empty")
	}

	_, errStyle := s.GetBeerStyles(beer.Style.Name)
	_, errType := s.GetBeerStyles(beer.Type.Name)

	// Verifica se ocorreu algum erro em ambas as chamadas
	if errStyle != nil || errType != nil {
		return errors.New("Style or Type not found")
	}

	// Verifica se os nomes de estilo e tipo não são vazios
	if beer.Style.Name == "" || beer.Type.Name == "" {
		return errors.New("Style or type cannot be empty")
	}

	return nil

}

func (s *BeerService) GetBeerStyles(b string) ([]*domain.BeerStyle, error) {
	return s.repository.GetBeerStyles(b)
}

func (s *BeerService) GetBeerTypes(b string) ([]*domain.BeerType, error) {
	return s.repository.GetBeerTypes(b)
}
