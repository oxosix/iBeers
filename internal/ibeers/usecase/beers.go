// usecase/beer_usecase.go
package usecase

import (
	"fmt"

	"github.com/d90ares/iBeers/internal/ibeers/domain"
)

type BeerUseCase struct {
	repository domain.Repository
}

func NewBeerUseCase(repository domain.Repository) *BeerUseCase {
	return &BeerUseCase{
		repository: repository,
	}
}

func (uc *BeerUseCase) GetAllBeers() ([]*domain.Beer, error) {
	// Implementar lógica para obter todas as cervejas usando o repositório
	fmt.Println("UseCase: Obter todas as cervejas")
	return nil, nil
}

// Implementar outros métodos de caso de uso, como GetBeerByID, SearchBeer, StoreBeer, etc.
