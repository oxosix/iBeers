// usecase/beer_usecase.go
package usecase

import (
	"fmt"

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

func (uc *BeerUseCase) GetAllBeers() ([]*domain.Beer, error) {
	// Implementar lógica para obter todas as cervejas usando o repositório
	fmt.Println("UseCase: Obter todas as cervejas")
	return nil, nil
}

// Implementar outros métodos de caso de uso, como GetBeerByID, SearchBeer, StoreBeer, etc.
